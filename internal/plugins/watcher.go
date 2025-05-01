package plugins

type Watcher interface {
	Start(cfg WatcherConfig, processor Processor, poster Poster)
}
