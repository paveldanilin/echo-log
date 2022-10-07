package service

import (
	"github.com/nxadm/tail"
)

type LineConsumer func(string)

type FileWatcher interface {
	Watch(c LineConsumer)
}

type tailFileWatcher struct {
	filename string
}

func NewTailFileWatcher(filename string) FileWatcher {
	return &tailFileWatcher{
		filename: filename,
	}
}

func (watcher *tailFileWatcher) Watch(c LineConsumer) {
	t, err := tail.TailFile(
		watcher.filename, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		c(line.Text)
	}
}
