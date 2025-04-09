package fwdq

import (
	"fmt"
	"time"
)

// FileName returns the file-name.
//
// A forward-queue (fwdq) is a date-time-based priority-queue.
// Where the 'priority' is 'time'.
// And, in particular, the closer a 'time' is to being the current time, without being in the past.
//
// Each element in a forward-queue (fwdq) are stored on a storate-drive in a file.
// The name of the file is a function of the 'time' (which is given by `when`) and the contents of the file (which is given by `bytes).`
//
// For example, if the `time` is given by the Unix-timestamp:
//
//	4822678189205111 // == 0x11223344556677
//
// and the contents of is:
//
//	[]byte{0xED,0x7B,0xA4,0x70,0x8E,0x54,0x46,0x5E,0x82,0x5C,0x99,0x71,0x20,0x43,0xE0,0x1C}
//
// Then the file-name returned will be:
//
//	"fwd_unixsec-0x0011223344556677_md5-0xa0a65c3a482379943f756a083e25f63f.octet-stream"
//
func FileName(when time.Time, bytes []byte) string {

	var encodedTime string = timeString(when)

	var encodedDigest string = digestString(bytes)

	return fmt.Sprintf("%s%s_%s%s", fileNamePrefix, encodedTime, encodedDigest, fileNameExtension)
}
