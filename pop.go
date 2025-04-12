package fwdq

import (
	"io/fs"
)

// Pop returns the element in the queue named `name`.
// The returned file needs to be closed.
func Pop(queuePath string, name string) (fs.File, error) {
	if "" == name {
		return nil, errEmptyFileName
	}
	if "" == queuePath {
		return nil, errEmptyQueuePath
	}

	return openInternalFile(queuePath, name)
}
