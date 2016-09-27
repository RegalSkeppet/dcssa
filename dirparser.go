package dcssa

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ParseDir recursively looks fror morgue files in the directory and parses them into data.
func ParseDir(dir string, data *Data) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			data.FailedReads[path] = err.Error()
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasPrefix(info.Name(), "morgue-") {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".txt") {
			return nil
		}
		ParseFile(path, data)
		return nil
	})
	if err != nil {
		return err
	}
	stats := NewStats()
	stats.Runs = len(data.Runs)
	racesWon := map[string]struct{}{}
	backgroundsWon := map[string]struct{}{}
	charactersWon := map[string]struct{}{}
	for _, run := range data.Runs {
		orb, ok := run.Attributes["0"]
		if !ok {
			continue
		}
		if len(orb) != 1 || orb[0] != "Orb of Zot" {
			continue
		}
		stats.Wins++
		racesWon[run.Race] = struct{}{}
		backgroundsWon[run.Background] = struct{}{}
		charactersWon[run.Race+" "+run.Background] = struct{}{}
	}
	for k := range racesWon {
		stats.RacesWon = append(stats.RacesWon, k)
	}
	sort.Strings(stats.RacesWon)
	for k := range backgroundsWon {
		stats.BackgroundsWon = append(stats.BackgroundsWon, k)
	}
	sort.Strings(stats.BackgroundsWon)
	for k := range charactersWon {
		stats.CharactersWon = append(stats.CharactersWon, k)
	}
	sort.Strings(stats.CharactersWon)
	data.Stats = stats
	return nil
}
