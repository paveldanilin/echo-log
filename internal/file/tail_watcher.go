package file

import (
	"github.com/nxadm/tail"
)

type TailMode int

const (
	TAIL_MODE_POLL    = 1
	TAIL_MODE_INOTIFY = 2
)

type TailWatcherConfig struct {
	Filename  string
	Follow    bool
	ReOpen    bool
	MustExist bool
	Mode      TailMode
}

type tailWatcher struct {
	config *TailWatcherConfig
	tail   *tail.Tail
}

func NewTailWatcher(config *TailWatcherConfig) Watcher {
	return &tailWatcher{
		config: config,
		tail:   nil,
	}
}

func (watcher *tailWatcher) Watch(c LineConsumer) error {
	defer func() {
		if watcher.tail != nil {
			watcher.tail.Cleanup()
			watcher.tail = nil
		}
	}()

	t, err := tail.TailFile(
		watcher.config.Filename, tail.Config{
			Follow:    watcher.config.Follow,
			ReOpen:    watcher.config.ReOpen,
			MustExist: watcher.config.MustExist,
			Poll:      watcher.config.Mode == TAIL_MODE_POLL, // for windows must be true
		})

	if err != nil {
		return err
	}

	watcher.tail = t

	for line := range t.Lines {
		c(line.Text)
	}

	return nil
}

func (watcher *tailWatcher) Stop() error {
	if watcher.tail != nil {
		return watcher.tail.Stop()
	}
	return nil
}
