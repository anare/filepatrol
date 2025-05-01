package plugins

type Auth interface {
	AddAuthHeaders(headers map[string]string, cfg WatcherConfig) error
}
