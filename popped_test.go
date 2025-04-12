package fwdq

import (
	"testing"
)

func TestPoppedFileNameMore(t *testing.T) {

	tests := []struct{
		FileName string
		UnixTime int64
		Chaos    uint64
		Expected string
	}{
		{
			Expected: "",
		},



		{
			FileName:                                                     "apple",
			Expected: "POPPED-0x0000000000000000-0x0000000000000000-POPPED-apple",
		},
		{
			FileName:                                                     "BANANA",
			UnixTime:         0x0123456789abcdef,
			Expected: "POPPED-0x0123456789abcdef-0x0000000000000000-POPPED-BANANA",
		},
		{
			FileName:        "Cherry",
			Chaos:                               0x0123456789abcdef,
			Expected: "POPPED-0x0000000000000000-0x0123456789abcdef-POPPED-Cherry",
		},
		{
			FileName:                                                     "dAtE",
			UnixTime:         0x0011223344556677,
			Chaos:                               0x0f1e2d3c4a5b6978,
			Expected: "POPPED-0x0011223344556677-0x0f1e2d3c4a5b6978-POPPED-dAtE",
		},



		{
			FileName:                                                     "fwd_unixsec-0x0011223344556677_md5-0xa0a65c3a482379943f756a083e25f63f.octet-stream",
			UnixTime:         0x000000000b49897c,
			Chaos:                               0x5678ea11b67dfc98,
			Expected: "POPPED-0x000000000b49897c-0x5678ea11b67dfc98-POPPED-fwd_unixsec-0x0011223344556677_md5-0xa0a65c3a482379943f756a083e25f63f.octet-stream",
		},
	}

	for testNumber, test := range tests {

		actual := poppedFileNameMore(test.FileName, test.UnixTime, test.Chaos)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual POPPED-file-name is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("FILE-NAME: %q", test.FileName)
			continue
		}
	}
}
