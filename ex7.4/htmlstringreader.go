package htmlstringreader

import (
	"bytes"
	"io"
)

// NewReader returns an io.Reader given a string
func NewReader(str string) io.Reader {
	buffer := bytes.NewReader([]byte(str))
	return buffer
}
