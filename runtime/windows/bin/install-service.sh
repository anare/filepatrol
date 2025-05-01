#!/bin/bash

SERVICE_NAME="filepatrol"
INSTALL_DIR="/opt/filepatrol"
BINARY_PATH="$INSTALL_DIR/filepatrol"

# Create install dir if not exists
sudo mkdir -p $INSTALL_DIR
sudo cp runtime/bin/linux/filepatrol $BINARY_PATH
sudo chmod +x $BINARY_PATH

# Create systemd service
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"

sudo bash -c "cat > $SERVICE_FILE" <<EOL
[Unit]
Description=FilePatrol Service
After=network.target

[Service]
ExecStart=$BINARY_PATH
WorkingDirectory=$INSTALL_DIR
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd and enable service
sudo systemctl daemon-reload
sudo systemctl enable $SERVICE_NAME
sudo systemctl start $SERVICE_NAME

echo "FilePatrol service installed and started."