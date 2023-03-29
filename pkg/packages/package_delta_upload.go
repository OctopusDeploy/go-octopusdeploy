package packages

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"github.com/OctopusDeploy/go-octodiff/pkg/octodiff"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type PackageSignatureResponse struct {
	Signature   string `json:"Signature" validate:"required"` // base64 encoded binary
	BaseVersion string `json:"BaseVersion" validate:"required"`
}

func fileBaseNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return filepath.Base(fileName[:pos])
	}
	return filepath.Base(fileName)
}

func attemptDeltaPush(
	client newclient.Client,
	spaceID string,
	fileName string,
	input io.ReadSeeker,
	inputLength int64,
	overwriteMode OverwriteMode) (outResponse *PackageUploadResponse, outCreatedNewPackage bool, outErr error) {
	// theoretically very old servers might not support delta push, but the Go Client is new enough that
	// we don't need to worry about gracefully handling those and can just fail with an HTTP 404 if we happen to hit one of those

	packageID, version, outErr := ParsePackageIDAndVersion(fileBaseNameWithoutExtension(fileName))
	if outErr != nil {
		return
	}

	params := map[string]any{
		"spaceId":   spaceID,
		"packageId": packageID,
		"version":   version,
	}
	signatureUri, outErr := client.URITemplateCache().Expand(uritemplates.PackageDeltaSignature, params)

	packageSignatureResponse, outErr := newclient.Get[PackageSignatureResponse](client.HttpSession(), signatureUri)
	if outErr != nil {
		// this can legitimately 404, which means the server doesn't have a prior version of this package and we should fallback to full upload
		// TODO do we let the caller handle the 404 or do we return some specific "package did not exist" indicator?
		return
	}
	signature, outErr := base64.StdEncoding.DecodeString(packageSignatureResponse.Signature)
	if outErr != nil {
		return
	}

	// Worst-case deltas for files can be equal to the size of the file itself. For something like a 5GB ISO, this means 5GB,
	// so we need to stream the delta into a temp file rather than hold it in memory.
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

	tmpFileInfo, outErr := tmpFile.Stat()
	if outErr != nil {
		return
	}
	ratio := float64(tmpFileInfo.Size()) / float64(inputLength)
	if ratio > 0.95 {
		// If the delta file is more than 95% the size of the full file, just upload the full file directly

		// TODO tell the caller to do a full upload. See-also 404 case above
		// TODO what about logging etc?
		return
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

	outResponse, outCreatedNewPackage, outErr = httpUploadPackageFile(client, deltaUploadUri, fileName, tmpFile)
	return
}
