package plugin_processor

import (
	"os"
	"path/filepath"

	"filepatrol/internal/plugins"
)

type SimpleProcessor struct{}

func (p *SimpleProcessor) Process(file plugins.FilePayload, poster plugins.Poster, processedDir string) error {
	if err := poster.Send(file, plugins.WatcherConfig{}); err != nil {
		return err
	}

	dest := filepath.Join(processedDir, filepath.Base(file.FilePath))

	// Ensure destination directory exists
	if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
		return err
	}

	err := os.Rename(file.FilePath, dest)
	if err != nil {
		return err
	}

	return nil
}

func NewProcessor() plugins.Processor {
	return &SimpleProcessor{}
}
