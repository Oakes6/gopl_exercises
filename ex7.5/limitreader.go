package limitreader

import (
	"io"
)

// type LimitReaderT struct {
// 	r io.Reader
// 	n int64
// }

// func LimitReader(r io.Reader, n int64) io.Reader {
// 	reader := LimitReaderT{r: r, n: n}
// 	return reader
// }

// func (lr LimitReaderT) Read(p []byte) (int, error) {
// 	limitedSizeSlice := make([]byte, lr.n, lr.n)
// 	num, e := lr.r.Read(limitedSizeSlice)
// 	if e != nil {
// 		return num, e
// 	}
// 	if int64(num) >= lr.n {
// 		return num, io.EOF
// 	}
// 	return num, nil
// }

type limitReader struct {
	underlyingReader io.Reader
	numStored        int64
	n                int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, 0, n}

}

func (lr *limitReader) Read(b []byte) (n int, err error) {
	lr.numStored += int64(len(b))
	if lr.numStored >= lr.n {
		b = b[:lr.numStored-lr.n]
		numRead, _ := lr.underlyingReader.Read(b)
		return numRead, io.EOF
	}

	return lr.underlyingReader.Read(b)
}
