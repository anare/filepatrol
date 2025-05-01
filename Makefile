PLUGIN_FLAGS=-buildmode=plugin

PLUGINS = \
	auth/jwt \
	poster/http \
	processor/simple \
	watcher/local_folder

cwd := $(shell pwd -L)
@echo Current dir: $(cwd);

# ==== Windows ====

build-win32:
	mkdir -p runtime/win32/bin runtime/win32/etc runtime/win32/cache runtime/win32/log
	GOOS=windows GOARCH=386 go build -o runtime/win32/bin/filepatrol.exe ./cmd/filepatrol

build-windows:
	mkdir -p runtime/windows/bin runtime/windows/etc runtime/windows/cache runtime/windows/log
	GOOS=windows GOARCH=amd64 go build -o runtime/windows/bin/filepatrol.exe ./cmd/filepatrol

# ==== macOS & Linux ====

build-mac: build-dir-mac build-binary-mac build-plugins-mac
build-linux: build-dir-linux build-binary-linux build-plugins-linux

build-dir-%:
	mkdir -p runtime/$*/bin runtime/$*/plugins runtime/$*/etc runtime/$*/cache runtime/$*/log

build-binary-mac:
	GOOS=darwin GOARCH=amd64 go build -o runtime/mac/bin/filepatrol ./cmd/filepatrol

build-binary-linux:
	GOOS=linux GOARCH=amd64 go build -o runtime/linux/bin/filepatrol ./cmd/filepatrol

build-plugins-mac:
	$(call build_plugins,mac,GOOS=darwin GOARCH=amd64)

build-plugins-linux:
	$(call build_plugins,linux,CGO_ENABLED=1 GOOS=linux GOARCH=amd64)

#		echo "🔧 Building plugin $$plugin for $(1)..."; \

define build_plugins
	@for plugin in $(PLUGINS); do \
		name=$$(basename $$plugin); \
		destdir=runtime/$(1)/plugins/$$(dirname $$plugin); \
		srcfile=plugins/$$plugin/main/$$name.go; \
		echo "🔧 Building plugin $$plugin for $(1) by path $$destdir of source file $$srcfile..."; \
		mkdir -p $$destdir; \
		$(2) go build $(PLUGIN_FLAGS) -o $$destdir/$$name.so $$srcfile || exit 1; \
	done
endef

# ==== Group Targets ====

build-windows-all: build-win32 build-windows
build-unix-all: build-mac build-linux
build-docker-all: build-win32 build-windows build-linux
build-all: build-windows-all build-unix-all

# ==== Clean Build ====

clean:
	@echo "🧹 Cleaning compiled binaries and plugins..."
	find runtime -type f \( -name 'filepatrol' -o -name 'filepatrol.exe' -o -name '*.so' \) -exec rm -v {} +

# ==== Plugin Structure Verifier ====

verify-plugins:
	go run cmd/tools/verify_plugins.go
