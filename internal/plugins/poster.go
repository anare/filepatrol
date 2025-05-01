package plugins

type Poster interface {
	Send(file FilePayload, cfg WatcherConfig) error
}
