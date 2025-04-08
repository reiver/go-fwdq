package fwdq

import (
	"fmt"
	"os"
	"math/rand"
	"time"

	"github.com/reiver/go-erorr"
	libpath "github.com/reiver/go-path"
)

func InsertBytes(queuePath string, when time.Time, bytes []byte) error {

	var temppath string
	var path string
	{
		var filename string = fileName(when, bytes)
		if "" == filename {
			return erorr.Errorf("fwdq: problem constructing file-name of fwdq element file using time %s and %d bytes — empty path", when, len(bytes))
		}

		path = libpath.Join(queuePath, filename)
		if "" == path {
			return erorr.Errorf("fwdq: problem constructing path to fwdq element file using queue-path %q and file-name %q — empty path", queuePath, filename)
		}

		var tempfilename string = fmt.Sprintf("spawn-%016x-%016x-%s", uint64(time.Now().Unix()), rand.Uint64(), filename)
		temppath = libpath.Join(queuePath, tempfilename)
		if "" == temppath {
			return erorr.Errorf("fwdq: problem constructing temp-path to fwdq element file using queue-path %q and temp-file-name %q — empty path", queuePath, tempfilename)
		}
	}

	// This is precautionary.
	// And, highly improbable to happen.
	//
	// If the spawn-*.fwd file already exists, we (try to) delete it.
	//
	// Since digests (from hash-functions) are involved, if the file already exist it SHOULD get the same content.
	//
	// We delete it just in case, for example, there was a crash before the file was fully written to.
	{
		err := os.Remove(temppath)
		if nil != err {
			
		}
	}

	{
		const perm os.FileMode = 0644

		err := os.WriteFile(temppath, bytes, perm)
		if nil != err {
			return erorr.Errorf("fwdq: problem writing fwdq element file (%q): %w", path, err)
		}
	}

	{
		err := os.Rename(temppath, path)
		if nil != err {
			return erorr.Errorf("fwdq: problem renaming temp-file %q to %q: %w", temppath, path, err)
		}
	}

	return nil
}
