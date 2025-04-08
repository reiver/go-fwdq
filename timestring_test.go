package fwdq

import (
	"testing"

	"time"
)

func TestTimeString(t *testing.T) {

	tests := []struct{
		Time time.Time
		Expected string
	}{
		{
			Time: time.Unix(0, 0),
			Expected: "unixsec-0x0000000000000000",
		},
		{
			Time: time.Unix(1, 0),
			Expected: "unixsec-0x0000000000000001",
		},
		{
			Time: time.Unix(2, 0),
			Expected: "unixsec-0x0000000000000002",
		},
		{
			Time: time.Unix(3, 0),
			Expected: "unixsec-0x0000000000000003",
		},
		{
			Time: time.Unix(4, 0),
			Expected: "unixsec-0x0000000000000004",
		},
		{
			Time: time.Unix(5, 0),
			Expected: "unixsec-0x0000000000000005",
		},
		{
			Time: time.Unix(6, 0),
			Expected: "unixsec-0x0000000000000006",
		},
		{
			Time: time.Unix(7, 0),
			Expected: "unixsec-0x0000000000000007",
		},
		{
			Time: time.Unix(8, 0),
			Expected: "unixsec-0x0000000000000008",
		},
		{
			Time: time.Unix(9, 0),
			Expected: "unixsec-0x0000000000000009",
		},
		{
			Time: time.Unix(10, 0),
			Expected: "unixsec-0x000000000000000a",
		},
		{
			Time: time.Unix(11, 0),
			Expected: "unixsec-0x000000000000000b",
		},
		{
			Time: time.Unix(12, 0),
			Expected: "unixsec-0x000000000000000c",
		},
		{
			Time: time.Unix(13, 0),
			Expected: "unixsec-0x000000000000000d",
		},
		{
			Time: time.Unix(14, 0),
			Expected: "unixsec-0x000000000000000e",
		},
		{
			Time: time.Unix(15, 0),
			Expected: "unixsec-0x000000000000000f",
		},
		{
			Time: time.Unix(16, 0),
			Expected: "unixsec-0x0000000000000010",
		},
		{
			Time: time.Unix(17, 0),
			Expected: "unixsec-0x0000000000000011",
		},
		{
			Time: time.Unix(18, 0),
			Expected: "unixsec-0x0000000000000012",
		},
		{
			Time: time.Unix(19, 0),
			Expected: "unixsec-0x0000000000000013",
		},
		{
			Time: time.Unix(20, 0),
			Expected: "unixsec-0x0000000000000014",
		},
		{
			Time: time.Unix(21, 0),
			Expected: "unixsec-0x0000000000000015",
		},
		{
			Time: time.Unix(22, 0),
			Expected: "unixsec-0x0000000000000016",
		},
		{
			Time: time.Unix(23, 0),
			Expected: "unixsec-0x0000000000000017",
		},
		{
			Time: time.Unix(24, 0),
			Expected: "unixsec-0x0000000000000018",
		},
		{
			Time: time.Unix(25, 0),
			Expected: "unixsec-0x0000000000000019",
		},
		{
			Time: time.Unix(26, 0),
			Expected: "unixsec-0x000000000000001a",
		},
		{
			Time: time.Unix(27, 0),
			Expected: "unixsec-0x000000000000001b",
		},
		{
			Time: time.Unix(28, 0),
			Expected: "unixsec-0x000000000000001c",
		},
		{
			Time: time.Unix(29, 0),
			Expected: "unixsec-0x000000000000001d",
		},
		{
			Time: time.Unix(30, 0),
			Expected: "unixsec-0x000000000000001e",
		},
		{
			Time: time.Unix(31, 0),
			Expected: "unixsec-0x000000000000001f",
		},
		{
			Time: time.Unix(32, 0),
			Expected: "unixsec-0x0000000000000020",
		},



		{
			Time: time.Unix(81985529216486895, 0),
			Expected: "unixsec-0x0123456789abcdef",
		},



		{
			Time: time.Unix(9141386507638288912, 0),
			Expected: "unixsec-0x7edcba9876543210",
		},
	}

	for testNumber, test := range tests {

		actual := timeString(test.Time)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual time-hex is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("TIME: %s", test.Time)
			continue
		}
	}
}
