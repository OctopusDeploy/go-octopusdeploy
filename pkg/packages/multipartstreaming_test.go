package packages_test

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestMultipartStreaming(t *testing.T) {
	fileContents := "The quick brown fox jumps over the lazy dogs"
	standardBoundary := "8082916287db208affca6f5f86fdca5fc8e34af10624000d7c821d303225" // easier if we don't have to sprintf boundaries in tests

	t.Run("small file where everything fits into a single segment", func(t *testing.T) {
		fileContents := "The quick brown fox jumps over the lazy dogs"
		m := packages.NewMultipartFileStreamingReader("foo.txt", strings.NewReader(fileContents))
		boundary := m.Boundary() // boundary will be a random value but always the same length
		// deliberately use random boundary here just as a sanity check we haven't broken things. All other tests can use standardBoundary

		buffer := make([]byte, 1024)
		cbWritten, err := m.Read(buffer)
		assert.Equal(t, io.EOF, err)
		assert.Equal(t, 283, cbWritten)
		assert.Equal(t, crlf(heredoc.Docf(`
		--%s
		Content-Disposition: form-data; name="file"; filename="foo.txt"
		Content-Type: application/octet-stream
		
		The quick brown fox jumps over the lazy dogs
		--%s--
		`, boundary, boundary)), string(buffer[:cbWritten]))
	})

	t.Run("medium file, header fits in one segment, rest fits in next segment", func(t *testing.T) {
		m := packages.NewMultipartFileStreamingReader("foo.txt", strings.NewReader(fileContents))
		_ = m.SetBoundary(standardBoundary)

		buffer := make([]byte, 188)

		// segment #1
		cbWritten, err := m.Read(buffer)
		assert.Equal(t, nil, err)
		assert.Equal(t, 188, cbWritten)
		assert.Equal(t, crlf(heredoc.Docf(`
		--8082916287db208affca6f5f86fdca5fc8e34af10624000d7c821d303225
		Content-Disposition: form-data; name="file"; filename="foo.txt"
		Content-Type: application/octet-stream
		
		The quick brown f`)), string(buffer[:cbWritten]))

		// segment #2
		cbWritten, err = m.Read(buffer)
		assert.Equal(t, io.EOF, err)
		assert.Equal(t, 95, cbWritten)
		assert.Equal(t, crlf(heredoc.Docf(`
		ox jumps over the lazy dogs
		--8082916287db208affca6f5f86fdca5fc8e34af10624000d7c821d303225--
		`)), string(buffer[:cbWritten]))
	})

	t.Run("large file, header fits in one segment, multiple body segments", func(t *testing.T) {
		longFileContents := strings.Repeat(fileContents, 7)
		m := packages.NewMultipartFileStreamingReader("foo.txt", strings.NewReader(longFileContents))
		_ = m.SetBoundary(standardBoundary)

		buffer := make([]byte, 188)

		// segment #1
		cbWritten, err := m.Read(buffer)
		assert.Equal(t, nil, err)
		assert.Equal(t, 188, cbWritten)
		assert.Equal(t, crlf(heredoc.Docf(`
		--8082916287db208affca6f5f86fdca5fc8e34af10624000d7c821d303225
		Content-Disposition: form-data; name="file"; filename="foo.txt"
		Content-Type: application/octet-stream
		
		The quick brown f`)), string(buffer[:cbWritten]))

		// segment #2
		cbWritten, err = m.Read(buffer)
		assert.Equal(t, nil, err)
		assert.Equal(t, 188, cbWritten)
		assert.Equal(t, "ox jumps over the lazy dogsThe quick brown fox jumps over the lazy dogsThe quick brown fox jumps over the lazy dogsThe quick brown fox jumps over the lazy dogsThe quick brown fox jumps ove", string(buffer[:cbWritten]))

		// segment #3
		cbWritten, err = m.Read(buffer)
		assert.Equal(t, io.EOF, err)
		assert.Equal(t, 171, cbWritten)
		assert.Equal(t, crlf(heredoc.Docf(`
		r the lazy dogsThe quick brown fox jumps over the lazy dogsThe quick brown fox jumps over the lazy dogs
		--8082916287db208affca6f5f86fdca5fc8e34af10624000d7c821d303225--
		`)), string(buffer[:cbWritten]))
	})

	t.Run("edge case; file fits in the initial buffer but the multipart trailer spills over", func(t *testing.T) {
		fileContents := "The quick brown fox jumps over the lazy dogs"
		m := packages.NewMultipartFileStreamingReader("foo.txt", strings.NewReader(fileContents))
		_ = m.SetBoundary(standardBoundary)

		buffer := make([]byte, 225)

		cbWritten, err := m.Read(buffer)
		assert.Equal(t, nil, err)
		assert.Equal(t, 225, cbWritten)
		assert.Equal(t, crlf(heredoc.Doc(`
		--8082916287db208affca6f5f86fdca5fc8e34af10624000d7c821d303225
		Content-Disposition: form-data; name="file"; filename="foo.txt"
		Content-Type: application/octet-stream
		
		The quick brown fox jumps over the lazy dogs
		--808291`)), string(buffer[:cbWritten]))

		cbWritten, err = m.Read(buffer)
		assert.Equal(t, io.EOF, err)
		assert.Equal(t, 58, cbWritten)
		assert.Equal(t, crlf(heredoc.Doc(`
		6287db208affca6f5f86fdca5fc8e34af10624000d7c821d303225--
		`)), string(buffer[:cbWritten]))
	})

	t.Run("edge case; catastrophic spill all over the place", func(t *testing.T) {
		fileContents := "The quick brown fox jumps over the lazy dogs"
		m := packages.NewMultipartFileStreamingReader("foo.txt", strings.NewReader(fileContents))
		_ = m.SetBoundary(standardBoundary)

		// use a large buffer to grab the full message, then we can assert each chunk in a loop without having to manually type them all
		largeBuffer := make([]byte, 1024)
		cbWritten, err := m.Read(largeBuffer)
		assert.Equal(t, io.EOF, err)
		expectedPayload := largeBuffer[:cbWritten]

		// ----

		m2 := packages.NewMultipartFileStreamingReader("foo.txt", strings.NewReader(fileContents))
		_ = m2.SetBoundary(standardBoundary)

		smallBuffer := make([]byte, 17) // this causes us to spill multiple times during the writing of the header, content, and trailer
		for i := 0; i < len(expectedPayload); i += len(smallBuffer) {
			expectedChunk := expectedPayload[i:min(i+len(smallBuffer), len(expectedPayload))]

			cbWritten, err := m2.Read(smallBuffer)
			if i+len(smallBuffer) >= len(expectedPayload) { // this is the last chunk
				assert.Equal(t, io.EOF, err)
			} else {
				assert.Equal(t, nil, err)
			}

			assert.Equal(t, len(expectedChunk), cbWritten)
			assert.Equal(t, string(expectedChunk), string(smallBuffer[:cbWritten])) // asserting on the stringified version gives us nicer error messages
		}
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func crlf(text string) string {
	return strings.ReplaceAll(text, "\n", "\r\n")
}
