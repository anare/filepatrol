package main

import (
	"filepatrol/internal/plugins"
	"filepatrol/plugins/watcher/local_folder/plugin"
)

// nolint:unused
func NewWatcher() plugins.Watcher {
	return &plugin_local_folder.LocalFolderWatcher{}
}
