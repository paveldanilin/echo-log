package service

import (
	"fmt"

	"github.com/nxadm/tail"
)

type FileWatcher interface {
	watch()
}

type tailFileWatcher struct {
	filename string
}

func (tailWatcher *tailFileWatcher) watch() {

	t, err := tail.TailFile(
		tailWatcher.filename, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		panic(err)
	}

	// Print the text of each received line
	for line := range t.Lines { // TODO: call processor
		fmt.Println(line.Text)
	}
}

/*
type cronFileWatcher struct {
	filename string
}
*/
