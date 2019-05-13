package main

import (
	"github.com/abonec/hugo_lastmod_updater/finder"
	"github.com/abonec/hugo_lastmod_updater/updater"
	"os"
)

const (
	defaultPath = "content/posts"
)

func main() {
	var path string
	if len(os.Args) == 1 {
		path = defaultPath
	} else {
		path = os.Args[1]
	}
	files, err := finder.NewFinder(path).Find()
	if err != nil {
		panic(err)
	}
	err = updater.NewUpdater(files).Update()
	if err != nil {
		panic(err)
	}
}
