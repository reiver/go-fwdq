package fwdq

import (
	"time"

	"github.com/reiver/go-hexadeca"
)

const timeStringPrefix string = "unixsec-0x"

const timeStringLength int = 16

func timeString(when time.Time) string {
	var unixtime uint64 = uint64(when.Unix())

	b15, b14, b13, b12, b11, b10, b9, b8, b7, b6, b5, b4, b3, b2, b1, b0 := hexadeca.EncodeUint64UsingLowerCaseSymbols(unixtime)
	var hex [timeStringLength]byte = [timeStringLength]byte{b15, b14, b13, b12, b11, b10, b9, b8, b7, b6, b5, b4, b3, b2, b1, b0}

	return timeStringPrefix + string(hex[:])
}
