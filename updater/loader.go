package updater

import "io/ioutil"

type loader interface {
	Load(path string) ([]byte, error)
}


type FileLoader struct {
}

func (FileLoader) Load(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
