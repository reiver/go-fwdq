package fwdq

import (
	"testing"

	"time"
)

func TestFileName(t *testing.T) {

	tests := []struct{
		Time time.Time
		Bytes []byte
		Expected string
	}{
		{
			Time: time.Unix(0, 0),
			Bytes: []byte(nil),
			Expected: "fwd_unixsec-0x0000000000000000_md5-0xd41d8cd98f00b204e9800998ecf8427e.octet-stream",
		},
		{
			Time: time.Unix(1, 0),
			Bytes: []byte{},
			Expected: "fwd_unixsec-0x0000000000000001_md5-0xd41d8cd98f00b204e9800998ecf8427e.octet-stream",
		},



		{
			Time: time.Unix(1, 0),
			Bytes: []byte{0x00},
			Expected: "fwd_unixsec-0x0000000000000001_md5-0x93b885adfe0da089cdf634904fd59f71.octet-stream",
		},
		{
			Time: time.Unix(2, 0),
			Bytes: []byte{0x01},
			Expected: "fwd_unixsec-0x0000000000000002_md5-0x55a54008ad1ba589aa210d2629c1df41.octet-stream",
		},
		{
			Time: time.Unix(3, 0),
			Bytes: []byte{0x02},
			Expected: "fwd_unixsec-0x0000000000000003_md5-0x9e688c58a5487b8eaf69c9e1005ad0bf.octet-stream",
		},
		{
			Time: time.Unix(4, 0),
			Bytes: []byte{0x03},
			Expected: "fwd_unixsec-0x0000000000000004_md5-0x8666683506aacd900bbd5a74ac4edf68.octet-stream",
		},
		{
			Time: time.Unix(5, 0),
			Bytes: []byte{0x04},
			Expected: "fwd_unixsec-0x0000000000000005_md5-0xec7f7e7bb43742ce868145f71d37b53c.octet-stream",
		},
		{
			Time: time.Unix(6, 0),
			Bytes: []byte{0x05},
			Expected: "fwd_unixsec-0x0000000000000006_md5-0x8bb6c17838643f9691cc6a4de6c51709.octet-stream",
		},
		{
			Time: time.Unix(7, 0),
			Bytes: []byte{0x06},
			Expected: "fwd_unixsec-0x0000000000000007_md5-0x06eca1b437c7904cc3ce6546c8110110.octet-stream",
		},
		{
			Time: time.Unix(8, 0),
			Bytes: []byte{0x07},
			Expected: "fwd_unixsec-0x0000000000000008_md5-0x89e74e640b8c46257a29de0616794d5d.octet-stream",
		},
		{
			Time: time.Unix(9, 0),
			Bytes: []byte{0x08},
			Expected: "fwd_unixsec-0x0000000000000009_md5-0xe2ba905bf306f46faca223d3cb20e2cf.octet-stream",
		},
		{
			Time: time.Unix(10, 0),
			Bytes: []byte{0x09},
			Expected: "fwd_unixsec-0x000000000000000a_md5-0x5e732a1878be2342dbfeff5fe3ca5aa3.octet-stream",
		},
		{
			Time: time.Unix(11, 0),
			Bytes: []byte{0x0A},
			Expected: "fwd_unixsec-0x000000000000000b_md5-0x68b329da9893e34099c7d8ad5cb9c940.octet-stream",
		},
		{
			Time: time.Unix(12, 0),
			Bytes: []byte{0x0B},
			Expected: "fwd_unixsec-0x000000000000000c_md5-0x13c8ffd977013703a701cf8e11deac65.octet-stream",
		},
		{
			Time: time.Unix(13, 0),
			Bytes: []byte{0x0C},
			Expected: "fwd_unixsec-0x000000000000000d_md5-0x58c89562f58fd276f592420068db8c09.octet-stream",
		},
		{
			Time: time.Unix(14, 0),
			Bytes: []byte{0x0D},
			Expected: "fwd_unixsec-0x000000000000000e_md5-0xdcb9be2f604e5df91deb9659bed4748d.octet-stream",
		},
		{
			Time: time.Unix(15, 0),
			Bytes: []byte{0x0E},
			Expected: "fwd_unixsec-0x000000000000000f_md5-0x4dedb2240a1e0f038dcdc8b3de92264c.octet-stream",
		},
		{
			Time: time.Unix(16, 0),
			Bytes: []byte{0x0F},
			Expected: "fwd_unixsec-0x0000000000000010_md5-0xd838691e5d4ad06879ca721442e883d4.octet-stream",
		},
		{
			Time: time.Unix(17, 0),
			Bytes: []byte{0x10},
			Expected: "fwd_unixsec-0x0000000000000011_md5-0x6b31bdfa7f9bfece263381ffa91bd6a9.octet-stream",
		},



		{
			Time: time.Unix(81985529216486895, 0),
			Bytes: []byte{0xFE},
			Expected: "fwd_unixsec-0x0123456789abcdef_md5-0x403ae091d3be6acf1181148527f1e0ae.octet-stream",
		},
		{
			Time: time.Unix(9141386507638288912, 0),
			Bytes: []byte{0xFF},
			Expected: "fwd_unixsec-0x7edcba9876543210_md5-0x00594fd4f42ba43fc1ca0427a0576295.octet-stream",
		},



		{
			Time: time.Unix(4822678189205111, 0),
			Bytes: []byte{0xED,0x7B,0xA4,0x70,0x8E,0x54,0x46,0x5E,0x82,0x5C,0x99,0x71,0x20,0x43,0xE0,0x1C},
			Expected: "fwd_unixsec-0x0011223344556677_md5-0xa0a65c3a482379943f756a083e25f63f.octet-stream",
		},
	}

	for testNumber, test := range tests {

		actual := fileName(test.Time, test.Bytes)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual time-hex is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("TIME: %#v", test.Time)
			t.Logf("BYTES: %#v", test.Bytes)
			continue
		}
	}
}
