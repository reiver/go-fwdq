package fwdq

import (
	"fmt"
	"math/rand"
	"time"
)

const poppedFileNameKeyWord string = "POPPED"

// When a file in the FWDQ directory is "popped", at first, it isn't actually deleted.
// But, is instead renamed.
//
// Basically, a "POPPED-%s-" prefix is added to the name.
//
// This will have an effect of making it invisible to the queue.
// While keep it around for a bit, just in case.
func poppedFileName(filename string) string {
	if "" == filename {
		return ""
	}

	var now int64 = time.Now().Unix()
	var chaos uint64 = rand.Uint64()

	return poppedFileNameMore(filename, now, chaos)
}

func poppedFileNameMore(filename string, unixtime int64, chaos uint64) string {
	if "" == filename {
		return ""
	}

	var prefix string = poppedFileNamePrefix(unixtime, chaos)
	return fmt.Sprintf("%s%s", prefix, filename)
}

func poppedFileNamePrefix(unixtime int64, chaos uint64) string {
	return fmt.Sprintf("%s-0x%016x-0x%016x-%s-", poppedFileNameKeyWord, uint64(unixtime), chaos, poppedFileNameKeyWord)
}
