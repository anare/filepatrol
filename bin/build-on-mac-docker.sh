#!/usr/bin/env bash

echo "Building the application..."
echo "Finding the root working directory..."
me="$(readlink "${BASH_SOURCE[0]}" || echo "${BASH_SOURCE[0]}")"
cwd=$(cd -P "$(dirname "${me}")" && pwd)
echo "Current working directory: $cwd"
if [ -f "${cwd}/../${me}" ]; then
  cwd=$(cd -P "$(dirname "$(dirname "$(readlink "${BASH_SOURCE[0]}" || echo "${BASH_SOURCE[0]}")")")" && pwd)
fi

echo "Use this as working directory: $cwd"

docker=$(which docker)
echo "Building for linux and win32 and windows using docker: ${docker}"
$docker compose -f docker/docker-compose.yml up --build
$docker compose -f docker/docker-compose.yml down
echo "Building for mac"
make build-mac
echo "Building the application completed."

source ./bin/zip.sh
