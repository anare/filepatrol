//go:build !windows

package patrol

import (
	loader "filepatrol/internal/loader"
	"filepatrol/internal/plugins"
)

func getWatcher(cfg plugins.WatcherConfig) plugins.Watcher {
	return loader.LoadWatcherPlugin(cfg.PluginPath)
}

func getProcessor(cfg plugins.WatcherConfig) plugins.Processor {
	return loader.LoadProcessorPlugin(cfg.PluginPath)
}

func getPoster(cfg plugins.WatcherConfig) plugins.Poster {
	return loader.LoadPostPlugin(cfg.PostPluginPath)
}
