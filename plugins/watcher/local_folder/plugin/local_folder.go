package plugin_local_folder

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"filepatrol/internal/plugins"
)

type LocalFolderWatcher struct{}

func (w *LocalFolderWatcher) Start(cfg plugins.WatcherConfig, processor plugins.Processor, poster plugins.Poster) {
	for {
		matches, err := filepath.Glob(filepath.Join(cfg.WatchDir, cfg.FileTemplate))
		if err != nil {
			log.Printf("Error watching dir: %v", err)
			continue
		}

		for _, match := range matches {
			content, err := os.ReadFile(match)
			if err != nil {
				log.Printf("Error reading file: %v", err)
				continue
			}

			payload := plugins.FilePayload{
				FilePath: match,
				Content:  content,
			}

			if err := processor.Process(payload, poster, cfg.ProcessedDir); err != nil {
				log.Printf("Processing error: %v", err)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
