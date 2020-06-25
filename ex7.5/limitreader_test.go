package limitreader

import (
	"bytes"
	"io"
	"testing"
)

func TestLimitReader(t *testing.T) {
	data := []byte{"parties"}
	seed := make([]byte, 8)
	bytesReader := bytes.Buffer{buf: data}
	limitReader := LimitReader(bytesReader, 6)
	n, e := limitReader.Read(seed)
	if e != io.EOF {
		t.Log("UNEXPECTED RESULT: ", n, " ", e)
		t.Fail()
	}
}
