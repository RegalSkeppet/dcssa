package dcssa

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
)

var reVersion = regexp.MustCompile(`Dungeon Crawl Stone Soup version ([\d\.]+)`)
var reWordPrefix = regexp.MustCompile(`^[^ ]+`)

// ParseFile parses a morgue file into data.
func ParseFile(path string, data *Data) error {
	run := &Run{}

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return errors.New("unexpected EOF")
	}

	run.Version = reVersion.FindStringSubmatch(scanner.Text())[1]

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

	data.Runs = append(data.Runs, run)
	return nil
}
