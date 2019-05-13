package updater

import (
	"io/ioutil"
	"os"
	"time"
)

type saver interface {
	Save(path string, content []byte, modTime time.Time, perm os.FileMode) error
}

type FileSaver struct {
}

func (FileSaver) Save(path string, content []byte, modTime time.Time, perm os.FileMode) error {
	err := ioutil.WriteFile(path, content, perm)
	if err != nil {
		return err
	}

	return os.Chtimes(path, time.Now(), modTime)
}
