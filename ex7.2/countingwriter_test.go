package countingwriter

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	p := []byte("thisishowweplaydestroy")
	buffer := &bytes.Buffer{}
	newWriter, size := CountingWriter(buffer)
	_, err := newWriter.Write(p)
	if err != nil {
		t.Log("Error: ", err)
		t.Fail()
	}
	p = []byte("sabbathbloddysabbath")
	_, err = newWriter.Write(p)
	if err != nil {
		t.Log("Error: ", err)
		t.Fail()
	}

	if *size != int64(42) {
		t.Log("UNEXPECTED RESULT: ", *size)
		t.Fail()
	}
}
