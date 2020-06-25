package limitreader

import (
	"io"
)

type LimitReaderT struct {
	r io.Reader
	n int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	reader := LimitReaderT{r: r, n: n}
	return reader
}

func (lr LimitReaderT) Read(p []byte) (int, error) {
	limitedSizeSlice := make([]byte, lr.n, lr.n)
	num, e := lr.r.Read(limitedSizeSlice)
	if e != nil {
		return num, e
	}
	if int64(num) >= lr.n {
		return num, io.EOF
	}
	return num, nil
}
