package counter

import (
	"testing"
)

func TestLineCounter(t *testing.T) {
	lc := &LineCounter{}
	p := []byte("Hello\nDebra\nI'm\nCounting\nLines")
	_, e := lc.Write(p)
	if e != nil {
		t.Log("error: ", e)
		t.Fail()
	}
	if lc.lines != 5 {
		t.Log("UNEXPECTED RESULT")
		t.Fail()
	}
}

func TestWordCounter(t *testing.T) {
	wc := &WordCounter{}
	p := []byte("Hello\nDebra\nI'm\nCounting\nLines\nlove happiness peace")
	_, e := wc.Write(p)
	if e != nil {
		t.Log("error: ", e)
		t.Fail()
	}
	if wc.words != 8 {
		t.Log("UNEXPECTED RESULT: ", wc.words)
		t.Fail()
	}
}
