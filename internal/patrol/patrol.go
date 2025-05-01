package patrol

import (
	"encoding/json"
	"log"
	"os"

	"filepatrol/internal/plugins"
)

var running = true

func Start(configPath string) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	var configs []plugins.WatcherConfig
	if err := json.Unmarshal(configData, &configs); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	for _, cfg := range configs {
		go startWatcher(cfg)
	}
}

func Stop() {
	running = false
}

func startWatcher(cfg plugins.WatcherConfig) {
	watcher := getWatcher(cfg)
	processor := getProcessor(cfg)
	poster := getPoster(cfg)

	watcher.Start(cfg, processor, poster)
}
