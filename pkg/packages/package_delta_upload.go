package packages

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/OctopusDeploy/go-octodiff/pkg/octodiff"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type PackageSignatureResponse struct {
	Signature   string `json:"Signature" validate:"required"` // base64 encoded binary
	BaseVersion string `json:"BaseVersion" validate:"required"`
}

func uploadDelta(
	client newclient.Client,
	spaceID string,
	fileName string,
	input io.ReadSeeker,
	inputLength int64,
	overwriteMode OverwriteMode) (outResponse *PackageUploadResponseV2, outErr error) {
	// theoretically very old servers might not support delta push, but the Go Client is new enough that
	// we don't need to worry about gracefully handling those and can just fail with an HTTP 404 if we happen to hit one of those

	packageID, version, outErr := ParsePackageIDAndVersion(filepath.Base(fileName))
	if outErr != nil {
		return
	}

	params := map[string]any{
		"spaceId":   spaceID,
		"packageId": packageID,
		"version":   version,
	}
	requestSignatureStartTime := time.Now()
	packageSignatureResponse, outErr := requestDeltaSignature(client, params)
	if outErr != nil {
		return
	}
	requestSignatureDuration := time.Since(requestSignatureStartTime)
	if packageSignatureResponse == nil {
		return deltaFallbackUploadFullFile(client, spaceID,
			fileName,
			input,
			inputLength,
			overwriteMode,
			0,
			requestSignatureDuration,
			0,
			DeltaBehaviourNoPreviousFile)
	}

	signature, outErr := base64.StdEncoding.DecodeString(packageSignatureResponse.Signature)
	if outErr != nil {
		return
	}

	// Worst-case deltas for files can be equal to the size of the file itself. For something like a 5GB ISO, this means 5GB,
	// so we need to stream the delta into a temp file rather than hold it in memory.
	buildDeltaStartTime := time.Now()
	tmpFile, outErr := os.CreateTemp("", "go-octopusdeploy-pkg-delta")
	if outErr != nil {
		return
	}
	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()
	bufferedTmpFileWriter := bufio.NewWriter(tmpFile)
	outErr = octodiff.NewDeltaBuilder().Build(input, inputLength, bytes.NewReader(signature), int64(len(signature)), octodiff.NewBinaryDeltaWriter(bufferedTmpFileWriter))
	if outErr != nil {
		return
	}
	outErr = bufferedTmpFileWriter.Flush()
	if outErr != nil {
		return
	}
	buildDeltaDuration := time.Since(buildDeltaStartTime)

	tmpFileInfo, outErr := tmpFile.Stat()
	if outErr != nil {
		return
	}
	ratio := float64(tmpFileInfo.Size()) / float64(inputLength)
	// If the delta file is more than 95% the size of the full file, just upload the full file directly
	if ratio > 0.95 {
		return deltaFallbackUploadFullFile(client, spaceID,
			fileName,
			input,
			inputLength,
			overwriteMode,
			tmpFileInfo.Size(),
			requestSignatureDuration,
			buildDeltaDuration,
			DeltaBehaviourNotEfficient)
	}

	_, outErr = tmpFile.Seek(0, io.SeekStart)
	if outErr != nil {
		return
	}

	// Do the actual delta upload here
	// note re-using params; we already have spaceId and packageId
	params["baseVersion"] = packageSignatureResponse.BaseVersion
	// Old C# client has logic for an `replace` parameter instead of `overwriteMode` but we choose not to support servers that old;
	// detecting it would require sniffing the links collection and we don't want to do that going forward
	if overwriteMode != "" {
		params["overwriteMode"] = overwriteMode
	}
	deltaUploadUri, outErr := client.URITemplateCache().Expand(uritemplates.PackageDeltaUpload, params)
	if outErr != nil {
		return
	}

	uploadStartTime := time.Now()
	stdResponse, createdNewPackage, outErr := httpUploadPackageFile(client, deltaUploadUri, fileName, tmpFile)
	if outErr != nil {
		return
	}
	uploadDuration := time.Since(uploadStartTime)

	return &PackageUploadResponseV2{
		CreatedNewFile:        createdNewPackage,
		UploadMethod:          UploadMethodDelta,
		PackageUploadResponse: *stdResponse,
		UploadInfo: &DeltaUploadedPackageInfo{
			FileSize:                 inputLength,
			DeltaSize:                tmpFileInfo.Size(),
			RequestSignatureDuration: requestSignatureDuration,
			BuildDeltaDuration:       buildDeltaDuration,
			UploadDuration:           uploadDuration,
			DeltaBehaviour:           DeltaBehaviourUploadedDeltaFile,
		},
	}, nil
}

func deltaFallbackUploadFullFile(
	client newclient.Client,
	spaceID string,
	fileName string,
	input io.ReadSeeker,
	inputLength int64,
	overwriteMode OverwriteMode,
	deltaLength int64,
	requestSignatureDuration time.Duration,
	buildDeltaDuration time.Duration,
	deltaBehaviour DeltaBehaviour) (*PackageUploadResponseV2, error) {
	// we can recover by pushing the full file
	// need to seek back to the start or the regular forward-read will fail
	_, err := input.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	fuParams := map[string]any{
		"spaceId": spaceID,
	}
	if overwriteMode != "" {
		fuParams["overwriteMode"] = overwriteMode
	}

	fullUploadUri, err := client.URITemplateCache().Expand(uritemplates.PackageUpload, fuParams)
	if err != nil {
		return nil, err
	}

	uploadStartTime := time.Now()
	stdResponse, createdNewPackage, err := httpUploadPackageFile(client, fullUploadUri, fileName, input)
	if err != nil {
		return nil, err
	}
	uploadDuration := time.Since(uploadStartTime)

	return &PackageUploadResponseV2{
		CreatedNewFile:        createdNewPackage,
		UploadMethod:          UploadMethodDelta,
		PackageUploadResponse: *stdResponse,
		UploadInfo: &DeltaUploadedPackageInfo{
			FileSize:                 inputLength,
			DeltaSize:                deltaLength,
			RequestSignatureDuration: requestSignatureDuration,
			BuildDeltaDuration:       buildDeltaDuration,
			UploadDuration:           uploadDuration,
			DeltaBehaviour:           deltaBehaviour,
		},
	}, nil
}

// requestDeltaSignature asks the server for a signature for a package (packageId and version in the params map).
// If the server returns 404 (not found) this indicates there's no existing package to delta off, in which case
// requestDeltaSignature will return (nil, nil) indiciating no signature but also no error.
func requestDeltaSignature(client newclient.Client, params map[string]any) (*PackageSignatureResponse, error) {
	signatureUri, err := client.URITemplateCache().Expand(uritemplates.PackageDeltaSignature, params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, signatureUri, nil)
	if err != nil {
		return nil, err
	}
	resp, outErr := client.HttpSession().DoRawRequest(req)
	if outErr != nil { // lower level error e.g. socket error, nonrecoverable
		return nil, err
	}
	defer newclient.CloseResponse(resp)

	var responseBody = new(PackageSignatureResponse)
	var responseError = new(core.APIError)
	bodyDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err = bodyDecoder.Decode(responseBody)
		if err != nil { // the server sent us a good response but we couldn't deserialize?
			return nil, err
		}
		return responseBody, nil
	} else {
		// The server responds with 404 if we ask for a signature and there are no prior versions of that package.
		// Give up attempting to delta and recover by just uploading the whole file.
		// Other kinds of errors are non-recoverable
		if resp.StatusCode == 404 {
			return nil, nil
		}

		// don't use core.APIErrorChecker, it's overly helpful and gets in the way of error handling.
		err = bodyDecoder.Decode(responseError)
		if err != nil { // can't deserialize the error JSON?
			return nil, err
		}
		// always return the error here, even if there was nothing to deserialize
		return nil, responseError
	}
}
