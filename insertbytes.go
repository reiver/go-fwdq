package fwdq

import (
	"fmt"
	"os"
	"math/rand"
	"time"

	"github.com/reiver/go-erorr"
	libpath "github.com/reiver/go-path"
)

// InsertBytes inserts a new element into the forward-queue (FWDQ).
//
// A forward-queue (FWDQ) is a storage-drive based queue, where the elements of the queue are files in a particular directory.
// The directory that the forward-queue (FWDQ) is in is given by `queuePath`.
//
// A forward-queue (FWDQ) is a priority-queue where the 'priority' is a date-time.
// That date-time is given by `when`.
//
// The content of the new element to insert into the forward-queue (FWDQ) is given by `bytes`.
//
// Example:
//
//	const queuePath string = "/path/to/queue"
//	
//	var when time.Time = time.Now().Add(5 * time.Hour)
//	
//	var bytes []byte = []byte("Hello world!")
//	
//	err := fwdw.InsertBytes(queuePath, when, bytes)
func InsertBytes(queuePath string, when time.Time, bytes []byte) error {

	var temppath string
	var path string
	{
		var filename string = FileName(when, bytes)
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
	// If the spawn-* file already exists, we (try to) delete it.
	//
	// Since digests (from hash-functions) are involved, if the file already exist it SHOULD get the same content.
	//
	// We delete it just in case, for example, there was a crash before the file was fully written to.
	//
	// Note that although we cannot guarantee a spawn-* file was fully written to, we can guarantee a regular file is fully written to,
	// since a regular file is a fully written to spawn-* that has been renamed to a regular file.
	{
		err := os.Remove(temppath)
		if nil != err {
			// We ignore the error.
//@TODO: we should only ignore the error if it is due to the file not existing (but pay attention to other errors).
			
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
