package fwdq

import (
	"testing"
)

func TestDigestString(t *testing.T) {

	tests := []struct{
		Bytes []byte
		Expected string
	}{
		{
			Bytes: []byte(nil),
			Expected: "md5-0xd41d8cd98f00b204e9800998ecf8427e",
		},
		{
			Bytes: []byte{},
			Expected: "md5-0xd41d8cd98f00b204e9800998ecf8427e",
		},



		{
			Bytes: []byte{0x00},
			Expected: "md5-0x93b885adfe0da089cdf634904fd59f71",
		},
		{
			Bytes: []byte{0x01},
			Expected: "md5-0x55a54008ad1ba589aa210d2629c1df41",
		},
		{
			Bytes: []byte{0x02},
			Expected: "md5-0x9e688c58a5487b8eaf69c9e1005ad0bf",
		},
		{
			Bytes: []byte{0x03},
			Expected: "md5-0x8666683506aacd900bbd5a74ac4edf68",
		},
		{
			Bytes: []byte{0x04},
			Expected: "md5-0xec7f7e7bb43742ce868145f71d37b53c",
		},
		{
			Bytes: []byte{0x05},
			Expected: "md5-0x8bb6c17838643f9691cc6a4de6c51709",
		},
		{
			Bytes: []byte{0x06},
			Expected: "md5-0x06eca1b437c7904cc3ce6546c8110110",
		},
		{
			Bytes: []byte{0x07},
			Expected: "md5-0x89e74e640b8c46257a29de0616794d5d",
		},
		{
			Bytes: []byte{0x08},
			Expected: "md5-0xe2ba905bf306f46faca223d3cb20e2cf",
		},
		{
			Bytes: []byte{0x09},
			Expected: "md5-0x5e732a1878be2342dbfeff5fe3ca5aa3",
		},
		{
			Bytes: []byte{0x0A},
			Expected: "md5-0x68b329da9893e34099c7d8ad5cb9c940",
		},
		{
			Bytes: []byte{0x0B},
			Expected: "md5-0x13c8ffd977013703a701cf8e11deac65",
		},
		{
			Bytes: []byte{0x0C},
			Expected: "md5-0x58c89562f58fd276f592420068db8c09",
		},
		{
			Bytes: []byte{0x0D},
			Expected: "md5-0xdcb9be2f604e5df91deb9659bed4748d",
		},
		{
			Bytes: []byte{0x0E},
			Expected: "md5-0x4dedb2240a1e0f038dcdc8b3de92264c",
		},
		{
			Bytes: []byte{0x0F},
			Expected: "md5-0xd838691e5d4ad06879ca721442e883d4",
		},
		{
			Bytes: []byte{0x10},
			Expected: "md5-0x6b31bdfa7f9bfece263381ffa91bd6a9",
		},



		{
			Bytes: []byte{0xFE},
			Expected: "md5-0x403ae091d3be6acf1181148527f1e0ae",
		},
		{
			Bytes: []byte{0xFF},
			Expected: "md5-0x00594fd4f42ba43fc1ca0427a0576295",
		},



		{
			Bytes: []byte{0xED,0x7B,0xA4,0x70,0x8E,0x54,0x46,0x5E,0x82,0x5C,0x99,0x71,0x20,0x43,0xE0,0x1C},
			Expected: "md5-0xa0a65c3a482379943f756a083e25f63f",
		},
	}

	for testNumber, test := range tests {

		actual := digestString(test.Bytes)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual time-hex is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("BYTES: %#v", test.Bytes)
			continue
		}
	}
}
