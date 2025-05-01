//go:build windows

package patrol

import (
	"filepatrol/internal/plugins"
	plugin_poster "filepatrol/plugins/poster/http/plugin"
	plugin_processor "filepatrol/plugins/processor/simple/plugin"
	plugin_local_folder "filepatrol/plugins/watcher/local_folder/plugin"
)

func getWatcher(cfg plugins.WatcherConfig) plugins.Watcher {
	// Windows-specific watcher implementation
	switch cfg.PluginPath {
	case "jwt":
	default:
		// Load the local folder watcher plugin
		return &plugin_local_folder.LocalFolderWatcher{}
	}

	return nil
}

func getProcessor(cfg plugins.WatcherConfig) plugins.Processor {
	switch cfg.PluginPath {
	case "processor":
	default:
		// Load the local folder watcher plugin
		return &plugin_processor.SimpleProcessor{}
	}

	return nil
}

func getPoster(cfg plugins.WatcherConfig) plugins.Poster {
	switch cfg.PluginPath {
	case "processor":
	default:
		// Load the local folder watcher plugin
		return &plugin_poster.SimplePoster{}
	}

	return nil
}
