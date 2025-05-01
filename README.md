# 🛡 FilePatrol – Plugin-Driven File Monitoring Service in Go

**FilePatrol** is a modular, future-proof service written in Go that monitors files across different sources (local
implemented) and processes them using plugin-based logic.

## 🔧 Features

- ✅ Runs as a service (Windows, systemd, launchd) using [`kardianos/service`](https://github.com/kardianos/service)
- ✅ Config-driven from `etc/config.json`
- ✅ Modular plugin system:
    - Watcher plugin
    - Processor plugin
    - Post plugin
    - Authentication plugin
- ✅ JWT token caching per watcher (`runtime/cache/{watcher_id}.jwt`)
- ✅ Auto-refresh JWT when expired
- ✅ Auto-retry once if HTTP POST fails with 401 Unauthorized
- ✅ Glob-based file matching (`*.txt`, `*.json`, etc.)
- ✅ Cross-platform builds: Linux, Windows, macOS

## 🧩 Project Structure

```
filepatrol/
├── cmd/filepatrol/
├── internal/patrol/
├── internal/plugins/
├── watcher_plugins/local_folder/
├── processor_plugins/processor/
├── post_plugins/post/
├── auth_plugins/jwt/
├── runtime/
│   ├── etc/config.json
│   ├── cache/
│   ├── bin/
│   └── plugins/
├── install-service.sh
├── install-service.ps1
├── go.mod
└── Makefile
```

## 📦 How to Build

```bash
make build-linux
make build-windows
make build-mac
make all
make clean
```

## 🛠 Authentication & Token Management

- JWT tokens are cached under `runtime/cache/` with one file per watcher.
- Tokens persist across application restarts.
- If a cached token is expired, FilePatrol automatically requests a fresh one.
- If a POST request fails with 401 Unauthorized, FilePatrol refreshes the token and retries once.

## 📜 Config Example (`etc/config.json`)

```json
[
  {
    "id": "local_watcher_1",
    "watchDir": "./watched",
    "fileTemplate": "*.txt",
    "postURL": "http://localhost:8080/upload",
    "pluginPath": "./runtime/plugins/linux/local_folder.so",
    "processorPluginPath": "./runtime/plugins/linux/processor.so",
    "postPluginPath": "./runtime/plugins/linux/post.so",
    "authPluginPath": "./runtime/plugins/linux/jwt.so",
    "jwtToken": "",
    "authURL": "http://localhost:8080/auth/login",
    "authBodyTemplate": "{\"username\":\"{username}\",\"password\":\"{password}\"}",
    "authUsername": "admin",
    "authPassword": "secret",
    "authTokenJSONField": "access_token",
    "processedDir": "./processed"
  }
]
```
