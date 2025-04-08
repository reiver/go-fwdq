package fwdq

import (
	"strings"
	"time"

	"github.com/reiver/go-hexadeca"
)

func parseFileName(filename string) (when time.Time, ok bool) {

	var str string = filename

	{
		if !strings.HasPrefix(str, fileNamePrefix) {
			var nada time.Time
			return nada, false
		}

		str = str[len(fileNamePrefix):]
	}

	{
		if !strings.HasPrefix(str, timeStringPrefix) {
			var nada time.Time
			return nada, false
		}

		str = str[len(timeStringPrefix):]
	}

	if !strings.HasSuffix(str, fileNameExtension) {
		var nada time.Time
		return nada, false
	}

	{
		var length int = len(str)

		if length < timeStringLength {
			return time.Time{}, false
		}

		str = str[:timeStringLength]
	}

	var unixtime uint64
	{
		var symbol15 byte = str[0]
		var symbol14 byte = str[1]
		var symbol13 byte = str[2]
		var symbol12 byte = str[3]
		var symbol11 byte = str[4]
		var symbol10 byte = str[5]
		var symbol9  byte = str[6]
		var symbol8  byte = str[7]
		var symbol7  byte = str[8]
		var symbol6  byte = str[9]
		var symbol5  byte = str[10]
		var symbol4  byte = str[11]
		var symbol3  byte = str[12]
		var symbol2  byte = str[13]
		var symbol1  byte = str[14]
		var symbol0  byte = str[15]

		var err error
		unixtime, err = hexadeca.DecodeUint64UsingLowerCaseSymbols(
			symbol15,
			symbol14,
			symbol13,
			symbol12,
			symbol11,
			symbol10,
			symbol9,
			symbol8,
			symbol7,
			symbol6,
			symbol5,
			symbol4,
			symbol3,
			symbol2,
			symbol1,
			symbol0,
		)
		if nil != err {
			return time.Time{}, false
		}
	}

	return time.Unix(int64(unixtime), 0), true
}
