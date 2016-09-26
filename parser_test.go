package dcssa

import (
	"os"
	"testing"
)

func TestParseVersion(t *testing.T) {
	file, err := os.Open("./morgue-Octodad-20160922-205833.txt")
	if err != nil {
		t.Fatal(err)
	}
	parser := NewParser(file)
	version, err := parser.ParseVersion()
	if err != nil {
		t.Fatal(err)
	}
	if version != "0.18.1" {
		t.Fatal(version)
	}
}
