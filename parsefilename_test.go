package fwdq

import (
	"testing"

	"time"
)

func TestParseFileName(t *testing.T) {

	tests := []struct{
		FileName string
		ExpectedWhen time.Time
		ExpectedOK   bool
	}{
		{
			FileName:  "fwd_unixsec-0x0000000000000000_md5-0x86fb269d190d2c85f6e0468ceca42a20.octet-stream",
			ExpectedWhen: time.Unix(0x0000000000000000, 0),
			ExpectedOK: true,
		},
		{
			FileName:  "fwd_unixsec-0x0000000067f4f089_md5-0x86fb269d190d2c85f6e0468ceca42a20.octet-stream",
			ExpectedWhen: time.Unix(0x0000000067f4f089, 0),
			ExpectedOK: true,
		},
		{
			FileName:  "fwd_unixsec-0x0000000067f508d1_md5-0x86fb269d190d2c85f6e0468ceca42a20.octet-stream",
			ExpectedWhen: time.Unix(0x0000000067f508d1, 0),
			ExpectedOK: true,
		},



		{
			FileName:  "fwd_unixsec-0x0000000067f4f089_md5-0x86fb269d190d2c85f6e0468ceca42a20.bytes",
			ExpectedWhen: time.Time{},
			ExpectedOK: false,
		},
	}

	for testNumber, test := range tests {

		actualWhen, actualOK := parseFileName(test.FileName)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual 'ok' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("FILE-NAME: %q", test.FileName)
				continue
			}
		}

		{
			expected := test.ExpectedWhen
			actual   := actualWhen

			if expected != actual {
				t.Errorf("For test #%d, the actual 'when' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %s", expected)
				t.Logf("ACTUAL:   %s", actual)
				t.Logf("FILE-NAME: %q", test.FileName)
				continue
			}
		}
	}
}
