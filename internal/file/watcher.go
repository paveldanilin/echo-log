package file

type LineConsumer func(string)

type Watcher interface {
	Watch(c LineConsumer)
}
