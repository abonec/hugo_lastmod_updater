package updater

import (
	"github.com/abonec/hugo_lastmod_updater/fileinfo"
	"log"
)

var (
	actualLoader loader = FileLoader{}
	actualSaver  saver  = FileSaver{}
)

type Updater struct {
	files []fileinfo.Info
}

func NewUpdater(files []fileinfo.Info) Updater {
	return Updater{
		files: files,
	}
}

func (u Updater) Update() error {
	for _, file := range u.files {
		err, wasUpdated := u.updateOne(file)
		logUpdate(file, wasUpdated, err)
	}
	return nil
}

func logUpdate(info fileinfo.Info, wasUpdated bool, err error) {
	if err == nil {
		if wasUpdated {
			log.Printf("File %s, modt: %s\n", info.FileName, info.ModTime)
		}
		return
	}
		log.Printf("File %s, error: %s\n", info.FileName, info.ModTime)
}

func (u Updater) updateOne(info fileinfo.Info) (error, bool) {
	content, err := actualLoader.Load(info.FileName)
	if err != nil {
		return err, false
	}
	updatedContent, wasUpdated, err := updateLastMod(content, info.ModTime)
	if err != nil {
		return err, wasUpdated
	}
	if wasUpdated {
		return actualSaver.Save(info.FileName, updatedContent, info.ModTime, info.FileMode), wasUpdated
	}
	return nil, wasUpdated
}
