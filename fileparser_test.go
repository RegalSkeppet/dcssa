package dcssa

import (
	"reflect"
	"testing"
)

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
	if data.Runs[0].Race != "Octopode" {
		t.Fatal(data.Runs[0].Race)
	}
	if data.Runs[0].Background != "Conjurer" {
		t.Fatal(data.Runs[0].Background)
	}
	if data.Runs[0].Turns != 167562 {
		t.Fatal(data.Runs[0].Turns)
	}
	if data.Runs[0].Time != "20:20:00" {
		t.Fatal(data.Runs[0].Time)
	}
	if !data.Runs[0].Escaped {
		t.Fatal(data.Runs[0].Escaped)
	}
	if !data.Runs[0].Orb {
		t.Fatal(data.Runs[0].Orb)
	}
	if !reflect.DeepEqual(data.Runs[0].Runes, []string{
		"barnacled", "slimy", "silver", "golden", "iron", "obsidian", "icy", "bone",
		"abyssal", "demonic", "glowing", "magical", "fiery", "dark", "gossamer",
	}) {
		t.Fatal(data.Runs[0].Runes)
	}
}

func TestParseLongRace(t *testing.T) {
	data := &Data{}
	ParseFile("./morgue-Kerumpuism-20160925-214510.txt", data)
	if len(data.Runs) != 1 {
		t.Fatal(len(data.Runs))
	}
	if data.Runs[0].Race != "High Elf" {
		t.Fatal(data.Runs[0].Race)
	}
	if data.Runs[0].Background != "Air Elementalist" {
		t.Fatal(data.Runs[0].Background)
	}
}
