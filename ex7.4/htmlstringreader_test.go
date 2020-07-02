package htmlstringreader

import (
	"testing"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	htmlstring := "<html><div><h1>Tanner's page</h1><h2>Trevor's section</h2></div></html>"
	reader := NewReader(htmlstring)
	actual, err := html.Parse(reader)
	if err != nil {
		t.Log(actual)
		t.Fail()
	}
	if actual == nil {
		t.Fail()
	}
}
