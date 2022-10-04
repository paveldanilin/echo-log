package service

type FileWatcher interface {
	watch()
	watchBackground()
	pause()
	stop()
}

type tailFileWatcher struct {
	filename string
}

type cronFileWatcher struct {
	filename string
}
