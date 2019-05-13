package fileinfo

import (
	"os"
	"time"
)

type Info struct {
	FileName string
	ModTime  time.Time
	FileMode os.FileMode
}
