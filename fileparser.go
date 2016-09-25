package dcssa

import (
	"bufio"
	"errors"
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

	// Parse name and background.
	line := nextLine(scanner, 0)
	matches := reNameBackground.FindStringSubmatch(line)
	split := strings.SplitN(matches[1], "the", 2)
	run.Name = strings.Trim(split[0], " ")
	run.Title = strings.Trim(split[1], " ")
	matches = reTurns.FindStringSubmatch(line)
	run.Turns, err = strconv.Atoi(matches[1])
	if err != nil {
		return err
	}
	matches = reTime.FindStringSubmatch(line)
	run.Time = matches[1]

	data.Runs = append(data.Runs, run)
	return nil
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
