package main

import (
	"filepatrol/internal/plugins"
	plugin_poster "filepatrol/plugins/poster/http/plugin"
)

// nolint:unused
func NewPoster() plugins.Poster { return &plugin_poster.SimplePoster{} }
