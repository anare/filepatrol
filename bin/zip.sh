#!/usr/bin/env bash

echo "Archiving the application..."
#zip_rev=$(git rev-parse --short HEAD)
zip_path="release/"
echo "Archiving to release folder: ${zip_path}"
mkdir -p "${zip_path}"

echo ""
echo "Archiving linux"
cd runtime/linux/ || exit 1
zip -r "../../${zip_path}/filepatrol-linux-amd64.zip" .

echo ""
echo "Archiving win32"
cd ../../ || exit 1
cd runtime/win32/ || exit 1
zip -r "../../${zip_path}/filepatrol-win32.zip" .

echo ""
echo "Archiving windows"
cd ../../ || exit 1
cd runtime/windows/ || exit 1
 zip -r "../../${zip_path}/filepatrol-windows-amd64.zip" .

echo ""
echo "Archiving mac"
cd ../../ || exit 1
cd runtime/mac/ || exit 1
 zip -r "../../${zip_path}/filepatrol-mac-amd64.zip" .
