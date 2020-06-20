package countingwriter

import (
	"io"
)

// A CountWriter wraps an io.Writer and maintains an int64 representing the
// total number of bytes written to its underlying data stream
type CountWriter struct {
	wrappedWriter io.Writer
	size          int64
}

func (w *CountWriter) Write(p []byte) (int, error) {
	size, err := w.wrappedWriter.Write(p)
	if err != nil {
		return 0, nil
	}
	w.size += int64(size)
	return size, nil
}

// CountingWriter wraps the given writer with the type CountWriter and returns it along with a pointer
// to an int64 containing a count of the total number of bytes written to the underlying data stream
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := &CountWriter{
		wrappedWriter: w,
		size:          0,
	}
	return newWriter, &newWriter.size
}
