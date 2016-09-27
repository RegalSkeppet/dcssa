package dcssa

import (
	"os"
	"path/filepath"
	"strings"
)

// ParseDir recursively looks fror morgue files in the directory and parses them into data.
func ParseDir(dir string, data *Data) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
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
}
