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
	if data.Runs[0].Name != "Octodad" {
		t.Fatal(data.Runs[0].Name)
	}
	if data.Runs[0].Title != "Conqueror" {
		t.Fatal(data.Runs[0].Title)
	}
	if data.Runs[0].Turns != 167562 {
		t.Fatal(data.Runs[0].Turns)
	}
	if data.Runs[0].Time != "20:20:00" {
		t.Fatal(data.Runs[0].Time)
	}
}
