package limitreader

import (
	"bytes"
	"io"
	"testing"
)

func TestLimitReader(t *testing.T) {
	data := []byte("parties")
	bytesReader := &bytes.Buffer{}
	n, e := bytesReader.Write(data)
	limitReader := LimitReader(bytesReader, 6)

	seed := make([]byte, 8)
	n, e = limitReader.Read(seed)
	if e != io.EOF {
		t.Log("UNEXPECTED RESULT: ", n, " ", e)
		t.Fail()
	}

	data = []byte("part")
	bytesReader.Reset()
	n, e = bytesReader.Write(data)

	limitReader = LimitReader(bytesReader, 6)
	seed = make([]byte, 8)
	n, e = limitReader.Read(seed)
	if e != nil && n != 4 {
		t.Log("UNEXPECTED RESULT: ", e, " ", n)
		t.Fail()
	}
}
