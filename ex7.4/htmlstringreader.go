package htmlstringreader

import (
	"io"
)

type StringReader struct {
	Source string
}

// return io.Reader that reads from string input instead of byte slice
func NewReader(str string) io.Reader {
	return &StringReader{Source: str}
}

func (sr *StringReader) Read(b []byte) (n int, err error) {
	n = copy(b, sr.Source)
	sr.Source = sr.Source[n:]
	if len(sr.Source) == 0 {
		return 0, io.EOF
	}
	return
}
