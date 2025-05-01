package main

import (
	"filepatrol/internal/plugins"
	"filepatrol/plugins/processor/simple/plugin"
)

// nolint:unused
func NewProcessor() plugins.Processor {
	return &plugin_processor.SimpleProcessor{}
}
