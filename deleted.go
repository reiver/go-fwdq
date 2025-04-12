package fwdq

import (
	"fmt"
)

const deleteFileNamePrefix string = "DELETED-"

// When a file in the FWDQ directory is "deleted", at first, it isn't actually deleted.
// But, is instead renamed.
//
// Basically, a "DELETED-" prefix is added to the name.
//
// This will have an effect of making it invisible to the queue.
// While keep it around for a bit, just in case.
func deletedFileName(filename string) string {
	if "" == filename {
		return ""
	}
	return fmt.Sprintf("%s%s", deleteFileNamePrefix, filename)
}
