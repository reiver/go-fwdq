package fwdq

import (
	"io/fs"
	"os"

	"github.com/reiver/go-erorr"
	path "github.com/reiver/go-path"
)

type internalFile struct {
	poppath string
	deletepath string
	file *os.File
}

func openInternalFile(queuePath string, filename string) (*internalFile, error) {

        var filepath  string = path.Join(queuePath, filename)
        var poppath   string = path.Join(queuePath, poppedFileName(filename))
	var deletepath string = path.Join(queuePath, deletedFileName(filename))

	{
		err := os.Rename(filepath, poppath)
		if nil != err {
			return nil, erorr.Errorf("fwdq: problem renaming queue-element-file %q to pop-file %q: %w", filepath, poppath, err)
		}
	}

	var file *os.File
	{

		var err error
		file, err = os.Open(poppath)
		if nil != err {
			return nil, erorr.Errorf("fwdq: problem opening pop-file %q: %w", poppath, err)
		}
		if nil == file {
			return nil, errNilFile
		}
	}

	return &internalFile{
		poppath: poppath,
		deletepath: deletepath,
		file: file,
	}, nil
}

func (receiver *internalFile) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	{
		err := receiver.file.Close()
		if nil != err {
			return erorr.Errorf("fwdq: problem closing pop-file %q: %w", receiver.poppath, err)
		}
	}

	{
		err := os.Rename(receiver.poppath, receiver.deletepath)
		if nil != err {
			return erorr.Errorf("fwdq: problem renaming pop-file %q to delete-file %q: %w", receiver.poppath, receiver.deletepath, err)
		}
	}

	{
		receiver.file = nil
		receiver.poppath = ""
		receiver.deletepath = ""
	}

	return nil
}

func (receiver *internalFile) Read(dst []byte) (n int, err error) {
	if nil == receiver {
		var nada int
		return nada, errNilReceiver
	}

	if nil == receiver.file {
		var nada int
		return nada, errNilFile
	}

	return receiver.file.Read(dst)
}

func (receiver *internalFile) Stat() (fs.FileInfo, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	if nil == receiver.file {
		return nil, errNilFile
	}

	return receiver.file.Stat()
}
