package dcssa

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reVersion = regexp.MustCompile(`Dungeon Crawl Stone Soup version ([\d\.]+)`)
var reWordPrefix = regexp.MustCompile(`^[^ ]+`)
var reNameBackground = regexp.MustCompile(`^([^\(]+)\(([^\)]+)\)`)
var reTurns = regexp.MustCompile(`Turns: (\d+)`)
var reTime = regexp.MustCompile(`Time: ([\d:]+)`)

// ParseFile parses a morgue file into data.
func ParseFile(path string, data *Data) error {
	run := &Run{}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	// Parse version.
	if !scanner.Scan() {
		return errors.New("unexpected EOF")
	}
	run.Version = reVersion.FindStringSubmatch(scanner.Text())[1]

	// Parse score.
	for {
		if !scanner.Scan() {
			return errors.New("unexpected EOF")
		}
		word := reWordPrefix.FindString(scanner.Text())
		if word != "" {
			run.Score, err = strconv.Atoi(word)
			if err != nil {
				return err
			}
			break
		}
	}

	// Parse name, title, race, background, turns and time.
	line := nextLine(scanner, 0)
	matches := reNameBackground.FindStringSubmatch(line)
	split := strings.SplitN(matches[1], "the", 2)
	run.Name = strings.Trim(split[0], " ")
	run.Title = strings.Trim(split[1], " ")

	split = strings.Split(matches[2], " ")
	run.Race, run.Background = parseRaceAndBackground(split)

	matches = reTurns.FindStringSubmatch(line)
	run.Turns, err = strconv.Atoi(matches[1])
	if err != nil {
		return err
	}
	matches = reTime.FindStringSubmatch(line)
	run.Time = matches[1]

	// Parse orb, runes and escaped.
	properties := [][]string{}
	for {
		if !scanner.Scan() {
			return errors.New("unexpected EOF")
		}
		if len(scanner.Text()) > 0 && scanner.Text()[0] == '@' {
			break
		}
	}
	var currentProperty []string
	for {
		if !scanner.Scan() {
			return errors.New("unexpected EOF")
		}
		line := scanner.Text()
		if len(line) == 0 {
			if currentProperty != nil {
				properties = append(properties, currentProperty)
			}
			break
		}
		if line[1] == ':' {
			if currentProperty != nil {
				properties = append(properties, currentProperty)
			}
			currentProperty = []string{line[0:1]}
			line = line[3:]
		}
		split := strings.Split(line, ",")
		for i := 0; i < len(split); i++ {
			split[i] = strings.Trim(split[i], " ")
		}

		currentProperty = append(currentProperty, split...)
	}
	for _, p := range properties {
		if p[0] == "O" {
			run.Orb = true
		}
		if p[0] == "}" {
			run.Runes = p[1:]
		}
	}
	fmt.Println(properties)
	line = nextLine(scanner, 0)
	if line == "You escaped." {
		run.Escaped = true
	}

	data.Runs = append(data.Runs, run)
	return nil
}

func parseRaceAndBackground(words []string) (string, string) {
	raceWords := 1
	switch words[0] {
	case "Deep", "High", "Hill", "Vine":
		raceWords = 2
	}
	return strings.Join(words[:raceWords], " "), strings.Join(words[raceWords:], " ")
}

func nextLine(s *bufio.Scanner, requireIndentation int) string {
lineloop:
	for {
		if !s.Scan() {
			return ""
		}
		line := s.Text()
		if len(line) <= requireIndentation {
			continue
		}
		i := 0
		for ; i < requireIndentation; i++ {
			if line[i] != ' ' {
				continue lineloop
			}
		}
		if line[i] != ' ' {
			return line
		}
	}
}
