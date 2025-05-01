package plugin_jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"filepatrol/internal/plugins"
)

type JWTAuth struct{}

func (j *JWTAuth) AddAuthHeaders(headers map[string]string, cfg plugins.WatcherConfig) error {
	token, err := readOrRequestJWT(cfg)
	if err != nil {
		return err
	}
	headers["Authorization"] = "Bearer " + token
	return nil
}

func readOrRequestJWT(cfg plugins.WatcherConfig) (string, error) {
	cacheFile := strings.ReplaceAll(cfg.Auth.Cache, "{id}", cfg.ID)

	if data, err := os.ReadFile(cacheFile); err == nil {
		token := string(data)
		if isTokenValid(token) {
			return token, nil
		}
	}

	// If no valid cached token, request a new one
	token, err := requestNewJWT(cfg)
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(cacheFile, []byte(token), 0644); err != nil {
		return token, err
	}

	return token, nil
}

func isTokenValid(token string) bool {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return false
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	var payload map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return false
	}

	expVal, ok := payload["exp"]
	if !ok {
		return false
	}

	expFloat, ok := expVal.(float64)
	if !ok {
		return false
	}

	expTime := time.Unix(int64(expFloat), 0)
	now := time.Now()

	// Allow small safety margin (e.g., 60 seconds)
	if expTime.Before(now.Add(60 * time.Second)) {
		return false
	}

	return true
}

func requestNewJWT(cfg plugins.WatcherConfig) (string, error) {
	body := cfg.Auth.BodyTemplate
	body = strings.ReplaceAll(body, "{username}", cfg.Auth.Username)
	body = strings.ReplaceAll(body, "{password}", cfg.Auth.Password)

	req, err := http.NewRequest("POST", cfg.Auth.URL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth server returned status %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return "", err
	}

	field, ok := parsed[cfg.Auth.TokenJSONField]
	if !ok {
		return "", errors.New("token field not found in auth response")
	}

	token, ok := field.(string)
	if !ok {
		return "", errors.New("token field is not a string")
	}

	return token, nil
}
