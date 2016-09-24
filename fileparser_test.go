package dcssa

import "testing"

func TestParseFile(t *testing.T) {
	data := &Data{}
	ParseFile("./morgue-Octodad-20160922-205833.txt", data)
	if len(data.Runs) != 1 {
		t.Fatal(len(data.Runs))
	}
	if data.Runs[0].Version != "0.18.1" {
		t.Fatal(data.Runs[0].Version)
	}
	if data.Runs[0].Score != 10097192 {
		t.Fatal(data.Runs[0].Score)
	}
}
