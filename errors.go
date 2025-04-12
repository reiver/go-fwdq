package fwdq

import (
	"github.com/reiver/go-erorr"
)

const (
	errEmptyFileName  = erorr.Error("fwdq; empty file-name")
	errEmptyQueuePath = erorr.Error("fwdq: empty queue-path")
	errNilFile        = erorr.Error("fwdq: nil file")
	errNilReceiver    = erorr.Error("fwdq: nil receiver")
)
