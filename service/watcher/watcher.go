package watcher

// or reader?

type Watcher interface {
	watch()
	stop()
}

type tailWatcher struct {
	filename string
}

func NewTailWatcher(filename string) Watcher {
	return &tailWatcher{
		filename: filename,
	}
}

func (tail *tailWatcher) watch() {

}

func (tail *tailWatcher) stop() {

}
