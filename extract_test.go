package htmlextr

import (
	"testing"
)

func TestExtractString(t *testing.T) {
	result, err := extractUrl("http://www.december.com/html/tutor/hello.html")
	if err != nil {
		t.Errorf("Error occured: %v\n", err)
	}
}
