package dcssa

import (
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	data := NewData()
	ParseFile("./morgue-Octodad-20160922-205833.txt", data)
	if len(data.Runs) != 1 {
		t.Fatal(data.FailedReads)
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
	if !reflect.DeepEqual(data.Runs[0].Stats, map[string]string{
		"Health": "224/224",
		"Magic":  "72/72",
		"Gold":   "9871",
		"AC":     "20",
		"EV":     "35",
		"SH":     "26",
		"Str":    "11",
		"Int":    "43",
		"Dex":    "20",
		"XL":     "27",
		"God":    "Sif Muna [******]",
		"Spells": "13 memorised, 0 levels left",
	}) {
		t.Fatal(data.Runs[0].Stats)
	}
	if !reflect.DeepEqual(data.Runs[0].Resistances, map[string]string{
		"rFire":    "+ . .",
		"rCold":    "+ + .",
		"rNeg":     "+ + +",
		"rPois":    "∞",
		"rElec":    ".",
		"rCorr":    "+",
		"SustAt":   "+",
		"MR":       "+++++",
		"Stlth":    "++........",
		"SeeInvis": "+",
		"Gourm":    ".",
		"Faith":    ".",
		"Spirit":   ".",
		"Dismiss":  ".",
		"Reflect":  ".",
		"Harm":     ".",
	}) {
		t.Fatal(data.Runs[0].Resistances)
	}
	if !reflect.DeepEqual(data.Runs[0].Equipped, []string{
		`+4 demon trident (elec)`,
		`+2 large shield {AC+3}`,
		`+2 hat "Kochixzary" {rF+ rCorr Stlth+}`,
		`amulet "Zisoixt" {+Rage Int+6}`,
		`ring "Vieplyot" {+Blink rC+ MP+9 Int+4 Dex+2}`,
		`+4 ring of protection`,
		`ring of Qiewk {+Fly rPois rN+ MR+ Dex+3}`,
		`ring of wizardry`,
		`+4 ring of protection`,
		`ring "Xeck" {MR+ MP+9 Str+2 Int+3}`,
		`ring of sustain attributes`,
		`ring of Devotion {Fragile +Inv Int+4 SInv}`,
	}) {
		t.Fatal(data.Runs[0].Equipped)
	}
	expectedAttributes := map[string][]string{
		"@": []string{
			"lich-form",
			"deflect missiles",
			"almost entirely resistant to hostile enchantments",
			"unstealthy",
		},
		"A": []string{
			"almost no armour",
			"amphibious",
			"8 rings",
			"constrict 8",
			"camouflage 1",
			"gelatinous body 1",
		},
		"a": []string{
			"End Transformation",
			"Channel Energy",
			"Forget Spell",
			"Renounce Religion",
			"Evoke Blink",
			"Evoke Invisibility",
			"Evoke Flight",
		},
		"0": []string{"Orb of Zot"},
		"}": []string{
			"barnacled",
			"slimy",
			"silver",
			"golden",
			"iron",
			"obsidian",
			"icy",
			"bone",
			"abyssal",
			"demonic",
			"glowing",
			"magical",
			"fiery",
			"dark",
			"gossamer",
		},
	}
	if len(data.Runs[0].Attributes) != len(expectedAttributes) {
		t.Fatal(data.Runs[0].Attributes)
	}
	for k, v := range data.Runs[0].Attributes {
		ev, ok := expectedAttributes[k]
		if !ok {
			t.Fatal(data.Runs[0].Attributes)
		}
		if !reflect.DeepEqual(v, ev) {
			t.Fatal(data.Runs[0].Attributes)
		}
	}
	if !data.Runs[0].Escaped {
		t.Fatal(data.Runs[0].Escaped)
	}
	if !reflect.DeepEqual(data.Runs[0].Skills, []Skill{
		{
			Name:         "Fighting",
			Level:        27,
			LevelDecimal: 0,
			Training:     false,
		},
		{
			Name:         "Polearms",
			Level:        14,
			LevelDecimal: 6,
			Training:     false,
		},
		{
			Name:         "Dodging",
			Level:        27,
			LevelDecimal: 0,
			Training:     false,
		},
		{
			Name:         "Stealth",
			Level:        7,
			LevelDecimal: 1,
			Training:     false,
		},
		{
			Name:         "Shields",
			Level:        25,
			LevelDecimal: 0,
			Training:     false,
		},
		{
			Name:         "Spellcasting",
			Level:        27,
			LevelDecimal: 0,
			Training:     false,
		},
		{
			Name:         "Conjurations",
			Level:        20,
			LevelDecimal: 0,
			Training:     false,
		},
		{
			Name:         "Charms",
			Level:        9,
			LevelDecimal: 9,
			Training:     false,
		},
		{
			Name:         "Necromancy",
			Level:        13,
			LevelDecimal: 8,
			Training:     false,
		},
		{
			Name:         "Translocations",
			Level:        13,
			LevelDecimal: 6,
			Training:     false,
		},
		{
			Name:         "Transmutations",
			Level:        12,
			LevelDecimal: 3,
			Training:     false,
		},
		{
			Name:         "Fire Magic",
			Level:        18,
			LevelDecimal: 9,
			Training:     false,
		},
		{
			Name:         "Air Magic",
			Level:        4,
			LevelDecimal: 6,
			Training:     false,
		},
		{
			Name:         "Earth Magic",
			Level:        19,
			LevelDecimal: 0,
			Training:     false,
		},
		{
			Name:         "Invocations",
			Level:        14,
			LevelDecimal: 5,
			Training:     false,
		},
		{
			Name:         "Evocations",
			Level:        10,
			LevelDecimal: 2,
			Training:     true,
		},
	}) {
		t.Fatal(data.Runs[0].Skills)
	}
}

func TestParseLongRace(t *testing.T) {
	data := NewData()
	ParseFile("./morgue-Kerumpuism-20160925-214510.txt", data)
	if len(data.Runs) != 1 {
		t.Fatal(data.FailedReads)
	}
	if data.Runs[0].Race != "High Elf" {
		t.Fatal(data.Runs[0].Race)
	}
	if data.Runs[0].Background != "Air Elementalist" {
		t.Fatal(data.Runs[0].Background)
	}
}

func TestParseWeirdName(t *testing.T) {
	data := NewData()
	ParseFile("./morgue-123_-aSdf456-20160927-094632.txt", data)
	if len(data.Runs) != 1 {
		t.Fatal(data.FailedReads)
	}
	if data.Runs[0].Name != "123 _-.aSdf 456" {
		t.Fatal(data.Runs[0].Name)
	}
	if data.Runs[0].Title != "Shield-Bearer" {
		t.Fatal(data.Runs[0].Title)
	}
}

func TestParseFelid(t *testing.T) {
	data := NewData()
	ParseFile("./morgue-Catwoman-20160824-105629.txt", data)
	if len(data.Runs) != 1 {
		t.Fatal(data.FailedReads)
	}
	if data.Runs[0].Stats["Lives"] != "0" {
		t.Fatal(data.Runs[0].Stats)
	}
	if data.Runs[0].Stats["Deaths"] != "2" {
		t.Fatal(data.Runs[0].Stats)
	}
}

func TestParseSkills(t *testing.T) {
	data := NewData()
	ParseFile("./morgue-Codiohudgh-20160816-151847.txt", data)
	if len(data.Runs) != 1 {
		t.Fatal(data.FailedReads)
	}
}
