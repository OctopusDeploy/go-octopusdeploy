package packages

import (
	"errors"
	"io"
	"mime/multipart"
)

type MultipartFileStreamingReader struct {
	IndirectWriter  *indirectWriter   // must initialize this before using the struct
	MultipartWriter *multipart.Writer // must initialize this before using the struct
	FileName        string            // must initialize this before using the struct
	FileReader      io.Reader         // must initialize this before using the struct

	currentPart        io.Writer // if we've started a multipart part block, this holds it so we can resume
	spill              []byte    // if we had bytes that didn't fit within a packet, these spill to the next read
	wroteFinalBoundary bool      // if we spill during the final boundary writing we must remember not to write it a second time
}

func NewMultipartFileStreamingReader(fileName string, fileReader io.Reader) *MultipartFileStreamingReader {
	indirectWriter := &indirectWriter{}

	return &MultipartFileStreamingReader{
		IndirectWriter:  indirectWriter,
		MultipartWriter: multipart.NewWriter(indirectWriter), // creates the random boundary, but doesn't assign an internal writer; this moves around
		FileName:        fileName,
		FileReader:      fileReader,
	}
}

// Read is called by the go HTTP Client, when it wants more bytes to send over the network.
// internally we generate multipart header/boundary data, and write it, combined with the file contents.
func (m *MultipartFileStreamingReader) Read(p []byte) (int, error) {
	pWriter := byteSliceWriter{Buf: p} // it's called pWriter because it writes to p
	var err error = nil

	// spillover from previous packet
	if m.spill != nil {
		cbToWrite := min(len(m.spill), len(p))
		_, err := pWriter.Write(m.spill[:cbToWrite])
		if err != nil { // no EOF here
			return 0, err
		}
		if pWriter.Remaining() == 0 {
			m.spill = m.spill[cbToWrite:]
			return pWriter.Pos, nil
		}
		// else; all the remaining spill fit in this read block, proceed and append more bytes to it
		m.spill = nil
	}

	if m.currentPart == nil { // we haven't written the part header yet
		m.IndirectWriter.Current = &pWriter
		m.currentPart, err = m.MultipartWriter.CreateFormFile("file", m.FileName)
		if err != nil {
			return 0, err
		}

		if pWriter.Spill != nil {
			m.spill = append(m.spill, pWriter.Spill...)
			pWriter.Spill = nil
			// data in the spill buffer guarantees we can't fit anything else; bail out now
			return pWriter.Pos, nil
		}
	} else {
		// routine copying
		m.IndirectWriter.Current = &pWriter
	}
	// copy as many bytes as will fit in the buffer. Can't spill here because we calculate it
	_, err = io.CopyN(m.currentPart, m.FileReader, int64(pWriter.Remaining()))
	if err == io.EOF {
		if !m.wroteFinalBoundary {
			m.wroteFinalBoundary = true
			e2 := m.MultipartWriter.Close()
			if e2 != nil {
				return 0, e2
			}
			if pWriter.Spill != nil {
				m.spill = append(m.spill, pWriter.Spill...)
				pWriter.Spill = nil
				// note: if we return here, technically the next Read should EOF. We don't need explicit code for that,
				// because the next read will consume the spill, fallthrough to the "copy file bytes" part, which will
				// copy zero bytes and return EOF again.
				return pWriter.Pos, nil
			}
		}
		// let the EOF fall through to the final return
	} else if err != nil {
		return 0, err
	}

	// return how many bytes were written to p
	return pWriter.Pos, err // we must return EOF if we get given it
}

func (m *MultipartFileStreamingReader) FormDataContentType() string {
	return m.MultipartWriter.FormDataContentType()
}

func (m *MultipartFileStreamingReader) Boundary() string {
	return m.MultipartWriter.Boundary()
}

func (m *MultipartFileStreamingReader) SetBoundary(boundary string) error {
	return m.MultipartWriter.SetBoundary(boundary)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type byteSliceWriter struct {
	Buf []byte
	Pos int

	// If writes fit within Buf, then Spill will be nil.
	// If a write is too big for Buf, then Spill will be the remaining bytes that couldn't fit.
	// You can only spill once; it is up to the caller to check for spill and deal with it before writing a second time
	Spill []byte
}

// Remaining returns the number of bytes in the underlying buffer that have not yet been written to
func (b *byteSliceWriter) Remaining() int {
	return len(b.Buf) - b.Pos
}

func (b *byteSliceWriter) Write(p []byte) (n int, err error) {
	if b.Spill != nil {
		return 0, errors.New("internal error: must flush any spilled bytes from byteSliceWriter before writing more")
	}

	cbToWrite := 0
	if b.Remaining() < len(p) {
		// we can't write all the bytes, so we need to spill some into a temporary buffer
		cbToWrite = b.Remaining()
		b.Spill = p[cbToWrite:]
	} else {
		// we can write them all
		cbToWrite = len(p)
		b.Spill = nil
	}

	subSlice := b.Buf[b.Pos : b.Pos+cbToWrite]
	bytesCopied := copy(subSlice, p)
	b.Pos = b.Pos + bytesCopied

	// even though we only write `bytesCopied` number of bytes, the spill buffer captures the rest,
	// so from the perspective of the caller, we did actually consume them all; we have to be consistent with our lie
	return len(p), err
}

// indirectWriter wraps an io.Writer, and passes through commands from outer to inner.
// It exists so you can create a single writer, feed it to an external source, then
// move the underlying writer around underneath it to support segmenting data across multiple packets
type indirectWriter struct {
	Current io.Writer
}

func (m *indirectWriter) Write(p []byte) (n int, err error) {
	if m.Current == nil {
		return 0, errors.New("internal error: indirectWriter.Current was nil on call to Write")
	}
	return m.Current.Write(p)
}
