package plugin_poster

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	loader "filepatrol/internal/loader"
	"filepatrol/internal/plugins"
)

type SimplePoster struct{}

func (p *SimplePoster) Send(file plugins.FilePayload, cfg plugins.WatcherConfig) error {
	return sendWithRetry(file, cfg, false)
}

func sendWithRetry(file plugins.FilePayload, cfg plugins.WatcherConfig, retried bool) error {
	headers := make(map[string]string)

	auth := loader.LoadAuthPlugin(cfg.AuthPluginPath)
	if auth != nil {
		err := auth.AddAuthHeaders(headers, cfg)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest("POST", cfg.PostURL, bytes.NewReader(file.Content))
	if err != nil {
		return err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized && !retried {
		log.Println("Received 401 Unauthorized, refreshing token and retrying...")
		forceRefreshJWT(cfg)
		return sendWithRetry(file, cfg, true)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("HTTP POST failed with status %d", resp.StatusCode)
	}

	return nil
}

func forceRefreshJWT(cfg plugins.WatcherConfig) {
	cacheFile := filepath.Join("runtime", "cache", fmt.Sprintf("%s.jwt", cfg.ID))
	if err := os.Remove(cacheFile); err != nil {
		log.Printf("Warning: failed to remove cached token: %v", err)
	}
}
