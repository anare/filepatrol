package plugins

import (
	"log"
	"plugin"

	. "filepatrol/internal/plugins"
)

func LoadWatcherPlugin(path string) Watcher {
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatalf("Failed to load watcher plugin %s: %v", path, err)
	}
	sym, err := p.Lookup("NewWatcher")
	if err != nil {
		log.Fatalf("Watcher plugin missing NewWatcher: %v", err)
	}
	return sym.(func() Watcher)()
}

func LoadProcessorPlugin(path string) Processor {
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatalf("Failed to load processor plugin %s: %v", path, err)
	}
	sym, err := p.Lookup("NewProcessor")
	if err != nil {
		log.Fatalf("Processor plugin missing NewProcessor: %v", err)
	}
	return sym.(func() Processor)()
}

func LoadPostPlugin(path string) Poster {
	p, err := plugin.Open(path)
	if err != nil {
		log.Fatalf("Failed to load post plugin %s: %v", path, err)
	}
	sym, err := p.Lookup("NewPoster")
	if err != nil {
		log.Fatalf("Post plugin missing NewPoster: %v", err)
	}
	return sym.(func() Poster)()
}

func LoadAuthPlugin(path string) Auth {
	p, err := plugin.Open(path)
	if err != nil {
		log.Printf("Failed to load auth plugin %s: %v", path, err)
		return nil
	}
	sym, err := p.Lookup("NewAuth")
	if err != nil {
		log.Printf("Auth plugin missing NewAuth: %v", err)
		return nil
	}
	return sym.(func() Auth)()
}
