package finder

import (
	"github.com/abonec/hugo_lastmod_updater/fileinfo"
	"log"
	"os"
	"path/filepath"
)

type Finder struct {
	base          string
	extensionName string
}

func NewFinder(basePath string) Finder {
	return Finder{
		base: basePath, extensionName: ".md",
	}
}

func (f Finder) Find() ([]fileinfo.Info, error) {
	var result []fileinfo.Info

	err := filepath.Walk(f.base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}
		if filepath.Ext(path) == f.extensionName {
			result = append(result, fileinfo.Info{FileName: path, ModTime: info.ModTime(), FileMode: info.Mode()})
		}
		return nil
	})
	if err != nil {
		return result, err
	}

	return result, nil
}
