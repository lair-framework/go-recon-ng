package reconng

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("results.json")
	if err != nil {
		t.Fatal(err)
	}
	output, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	re, err := Parse(output)
	if err != nil {
		t.Fatal(err)
	}
	if re == nil {
		t.Fatal("output nil")
	}
}
