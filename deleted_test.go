package fwdq

import (
	"testing"
)

func TestDeletedFileName(t *testing.T) {

	tests := []struct{
		FileName string
		Expected string
	}{
		{
			FileName: "",
			Expected: "",
		},



		{
			FileName:         "apple",
			Expected: "DELETED-apple",
		},
		{
			FileName:         "BANANA",
			Expected: "DELETED-BANANA",
		},
		{
			FileName:         "Cherry",
			Expected: "DELETED-Cherry",
		},
		{
			FileName:         "dAtE",
			Expected: "DELETED-dAtE",
		},



		{
			FileName:         "fwd_unixsec-0x0011223344556677_md5-0xa0a65c3a482379943f756a083e25f63f.octet-stream",
			Expected: "DELETED-fwd_unixsec-0x0011223344556677_md5-0xa0a65c3a482379943f756a083e25f63f.octet-stream",
		},
	}

	for testNumber, test := range tests {

		actual := deletedFileName(test.FileName)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual DELETED-file-name is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("FILE-NAME: %q", test.FileName)
			continue
		}
	}
}
