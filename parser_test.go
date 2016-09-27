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
			State:        MASTERED,
		},
		{
			Name:         "Polearms",
			Level:        14,
			LevelDecimal: 6,
			State:        USING,
		},
		{
			Name:         "Dodging",
			Level:        27,
			LevelDecimal: 0,
			State:        MASTERED,
		},
		{
			Name:         "Stealth",
			Level:        7,
			LevelDecimal: 1,
			State:        USING,
		},
		{
			Name:         "Shields",
			Level:        25,
			LevelDecimal: 0,
			State:        USING,
		},
		{
			Name:         "Spellcasting",
			Level:        27,
			LevelDecimal: 0,
			State:        MASTERED,
		},
		{
			Name:         "Conjurations",
			Level:        20,
			LevelDecimal: 0,
			State:        USING,
		},
		{
			Name:         "Charms",
			Level:        9,
			LevelDecimal: 9,
			State:        USING,
		},
		{
			Name:         "Necromancy",
			Level:        13,
			LevelDecimal: 8,
			State:        USING,
		},
		{
			Name:         "Translocations",
			Level:        13,
			LevelDecimal: 6,
			State:        USING,
		},
		{
			Name:         "Transmutations",
			Level:        12,
			LevelDecimal: 3,
			State:        USING,
		},
		{
			Name:         "Fire Magic",
			Level:        18,
			LevelDecimal: 9,
			State:        USING,
		},
		{
			Name:         "Air Magic",
			Level:        4,
			LevelDecimal: 6,
			State:        USING,
		},
		{
			Name:         "Earth Magic",
			Level:        19,
			LevelDecimal: 0,
			State:        USING,
		},
		{
			Name:         "Invocations",
			Level:        14,
			LevelDecimal: 5,
			State:        USING,
		},
		{
			Name:         "Evocations",
			Level:        10,
			LevelDecimal: 2,
			State:        TRAINING,
		},
	}) {
		t.Fatal(data.Runs[0].Skills)
	}
	if !reflect.DeepEqual(data.Runs[0].Spells, []Spell{
		{
			Name:    "Lehudib's Crystal Sp",
			Type:    "Conj/Erth",
			Power:   "########..",
			Failure: 1,
			Level:   8,
			Hunger:  "N/A",
		},
		{
			Name:    "Fire Storm",
			Type:    "Conj/Fire",
			Power:   "########..",
			Failure: 4,
			Level:   9,
			Hunger:  "N/A",
		},
		{
			Name:    "Ozocubu's Armour",
			Type:    "Chrm/Ice",
			Power:   "#######.",
			Failure: 1,
			Level:   3,
			Hunger:  "N/A",
		},
		{
			Name:    "Controlled Blink",
			Type:    "Tloc",
			Power:   "N/A",
			Failure: 4,
			Level:   8,
			Hunger:  "N/A",
		},
		{
			Name:    "Deflect Missiles",
			Type:    "Chrm/Air",
			Power:   "#######...",
			Failure: 1,
			Level:   6,
			Hunger:  "N/A",
		},
		{
			Name:    "Spellforged Servitor",
			Type:    "Conj/Summ",
			Power:   "#######...",
			Failure: 1,
			Level:   7,
			Hunger:  "N/A",
		},
		{
			Name:    "Blink",
			Type:    "Tloc",
			Power:   "N/A",
			Failure: 0,
			Level:   2,
			Hunger:  "N/A",
		},
		{
			Name:    "Haste",
			Type:    "Chrm",
			Power:   "#######.",
			Failure: 1,
			Level:   6,
			Hunger:  "N/A",
		},
		{
			Name:    "Apportation",
			Type:    "Tloc",
			Power:   "########..",
			Failure: 0,
			Level:   1,
			Hunger:  "N/A",
		},
		{
			Name:    "Shatter",
			Type:    "Erth",
			Power:   "########..",
			Failure: 4,
			Level:   9,
			Hunger:  "N/A",
		},
		{
			Name:    "Disjunction",
			Type:    "Tloc",
			Power:   "########..",
			Failure: 4,
			Level:   8,
			Hunger:  "N/A",
		},
		{
			Name:    "Necromutation",
			Type:    "Trmt/Necr",
			Power:   "########..",
			Failure: 4,
			Level:   8,
			Hunger:  "N/A",
		},
		{
			Name:    "Irradiate",
			Type:    "Conj/Trmt",
			Power:   "########..",
			Failure: 0,
			Level:   5,
			Hunger:  "N/A",
		},
	}) {
		t.Fatal(data.Runs[0].Spells)
	}
	if !reflect.DeepEqual(data.Runs[0].Mutations, []string{
		"You cannot wear most types of armour.",
		"You are amphibious.",
		"You can wear up to eight rings at the same time.",
		"You can use your tentacles to constrict eight enemies at once.",
		"Your skin changes colour to match your surroundings (Stealth).",
		"Your rubbery body absorbs attacks (AC +1).",
	}) {
		t.Fatal(data.Runs[0].Mutations)
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
	if !reflect.DeepEqual(data.Runs[0].Skills, []Skill{
		{
			Name:         "Dodging",
			Level:        2,
			LevelDecimal: 0,
			State:        USING,
		},
		{
			Name:         "Stealth",
			Level:        3,
			LevelDecimal: 3,
			State:        TRAINING,
		},
		{
			Name:         "Spellcasting",
			Level:        2,
			LevelDecimal: 2,
			State:        TRAINING,
		},
		{
			Name:         "Transmutations",
			Level:        1,
			LevelDecimal: 0,
			State:        UNUSED,
		},
		{
			Name:         "Earth Magic",
			Level:        3,
			LevelDecimal: 3,
			State:        FOCUSING,
		},
	}) {
		t.Fatal(data.Runs[0].Skills)
	}
}
