package fwdq

import (
	"fmt"
	"time"
)

func fileName(when time.Time, bytes []byte) string {

	var encodedTime string = timeString(when)

	var encodedDigest string = digestString(bytes)

	return fmt.Sprintf("%s%s_%s%s", fileNamePrefix, encodedTime, encodedDigest, fileNameExtension)
}
