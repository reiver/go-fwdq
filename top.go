package fwdq

import (
	"io/fs"
	"path/filepath"
	"time"

	"github.com/reiver/go-erorr"
)

func Top(queuePath string, minTime time.Time) ([]string, error) {

	var matches []string

	var fn fs.WalkDirFunc = func(path string, entry fs.DirEntry, err error) error {
		if entry.IsDir() {
			return nil
		}

		var name string = entry.Name()

		when, ok := parseFileName(name)
		if !ok {
			return nil
		}

		if when.Before(minTime) {
			return nil
		}

		matches = append(matches, name)

		return nil
	}

	err := filepath.WalkDir(queuePath, fn)
	if nil != err {
		return nil, erorr.Errorf("fwdq: problem walking queue-directory %q: %w", queuePath, err)
	}

	return matches, nil
}
