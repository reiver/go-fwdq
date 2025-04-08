package fwdq

import (
	"crypto/md5"
	"encoding/hex"
)

const digestPrefix string = "md5-0x"

func digestString(bytes []byte) string {
	result := md5.Sum(bytes)
	return digestPrefix + hex.EncodeToString(result[:])
}
