package file

import "github.com/nxadm/tail"

type tailWatcher struct {
	filename string
}

func NewTailWatcher(filename string) Watcher {
	return &tailWatcher{
		filename: filename,
	}
}

func (watcher *tailWatcher) Watch(c LineConsumer) {
	t, err := tail.TailFile(
		watcher.filename, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		panic(err)
	}

	for line := range t.Lines {
		c(line.Text)
	}
}
