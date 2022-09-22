package packages

import (
	"encoding/json"
	"errors"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"io"
	"mime/multipart"
	"net/http"
)

type PackageService struct {
	bulkPath           string
	deltaSignaturePath string
	deltaUploadPath    string
	notesListPath      string
	uploadPath         string

	services.CanDeleteService
}

func NewPackageService(sling *sling.Sling, uriTemplate string, deltaSignaturePath string, deltaUploadPath string, notesListPath string, bulkPath string, uploadPath string) *PackageService {
	return &PackageService{
		bulkPath:           bulkPath,
		deltaSignaturePath: deltaSignaturePath,
		deltaUploadPath:    deltaUploadPath,
		notesListPath:      notesListPath,
		uploadPath:         uploadPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServicePackageService, sling, uriTemplate),
		},
	}
}

// GetAll returns all packages. If none can be found or an error occurs, it
// returns an empty collection.
func (s *PackageService) GetAll() ([]*Package, error) {
	items := []*Package{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return items, err
}

// GetByID returns the package that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s *PackageService) GetByID(id string) (*Package, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Package), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Package), nil
}

type multipartFileStreamingReader struct {
	MovableWriter   *MovableWriter    // must initialize this before using the struct
	MultipartWriter *multipart.Writer // must initialize this before using the struct
	FileName        string            // must initialize this before using the struct
	FileReader      io.Reader         // must initialize this before using the struct

	currentPart io.Writer // internal state tracking
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// surely go has this built-in, but I can't find it
type byteSliceWriter struct {
	Buf []byte
	Pos int
}

// Remaining returns the number of bytes in the underlying buffer that have not yet been written to
func (b *byteSliceWriter) Remaining() int {
	return len(b.Buf) - b.Pos
}

func (b *byteSliceWriter) Write(p []byte) (n int, err error) {
	cb := min(
		b.Remaining(),
		len(p))

	subSlice := b.Buf[b.Pos : b.Pos+cb]
	bytesCopied := copy(subSlice, p)
	b.Pos = b.Pos + bytesCopied

	return bytesCopied, err
}

type MovableWriter struct {
	CurrentWriter io.Writer
}

func (m *MovableWriter) Write(p []byte) (n int, err error) {
	if m.CurrentWriter == nil {
		return 0, errors.New("MovableWriter.CurrentWriter was nil on call to Write")
	}
	return m.CurrentWriter.Write(p)
}

// conforms to io.Reader
func (m *multipartFileStreamingReader) Read(p []byte) (int, error) {
	pWriter := byteSliceWriter{Buf: p}
	var err error = nil

	if m.currentPart == nil { // we haven't written the part header yet
		m.MovableWriter.CurrentWriter = &pWriter
		m.currentPart, err = m.MultipartWriter.CreateFormFile("file", m.FileName)
		// TODO: if there's not enough space in the buffer to write the initial multipart header we must spill over into a temp buffer
		// and let the other end call Read again to pick it up
		if err != nil {
			return 0, err
		}
	} else {
		// routine copying until we've done it all
		m.MovableWriter.CurrentWriter = &pWriter
	}
	// copy as many bytes as will fit in the buffer
	_, err = io.CopyN(m.currentPart, m.FileReader, int64(pWriter.Remaining()))
	if err == io.EOF {
		// writes the final boundary.
		// TODO: if there's not enough space remaining in the buffer we must spill over into a temp buffer
		// and let the other end call Read again to pick it up
		e2 := m.MultipartWriter.Close()
		if e2 != nil {
			return 0, e2
		}
	} else if err != nil {
		return 0, err
	}

	// return how many bytes were written to p
	return pWriter.Pos, err // we must return EOF if we get given it
}

func (m *multipartFileStreamingReader) FormDataContentType() string {
	return m.MultipartWriter.FormDataContentType()
}

func Upload(client newclient.Client, command *PackageUploadCommand) (*PackageUploadResponse, bool, error) {
	movableWriter := &MovableWriter{}
	m := &multipartFileStreamingReader{
		MovableWriter:   movableWriter,
		MultipartWriter: multipart.NewWriter(movableWriter), // creates the random boundary, but doesn't assign an internal writer; this moves around
		FileName:        command.FileName,
		FileReader:      command.FileReader,
	}

	path, err := client.URITemplateCache().Expand(uritemplates.PackageUpload, command)
	if err != nil {
		return nil, false, err
	}

	req, err := http.NewRequest(http.MethodPost, path, m)
	if err != nil {
		return nil, false, err
	}
	req.Header.Add("Content-Type", m.FormDataContentType())

	resp, err := client.HttpSession().DoRawRequest(req)
	if err != nil {
		return nil, false, err
	}
	defer newclient.CloseResponse(resp)

	bodyDecoder := json.NewDecoder(resp.Body)
	if resp.StatusCode == 201 || resp.StatusCode == 200 {
		outputResponseBody := new(PackageUploadResponse)
		err = bodyDecoder.Decode(outputResponseBody)
		if err != nil {
			return nil, false, err
		}
		// the server returns 201 if it created a new file, 200 if it ignored an existing file
		createdNewFile := resp.StatusCode == 201
		return outputResponseBody, createdNewFile, nil
	} else {
		outputResponseError := new(core.APIError)
		err = bodyDecoder.Decode(outputResponseError)
		if err != nil {
			return nil, false, err
		}
		return nil, false, outputResponseError
	}
}

// Update modifies a package based on the one provided as input.
func (s *PackageService) Update(octopusPackage *Package) (*Package, error) {
	if octopusPackage == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, "octopusPackage")
	}

	path, err := services.GetUpdatePath(s, octopusPackage)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), octopusPackage, new(Package), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Package), nil
}
