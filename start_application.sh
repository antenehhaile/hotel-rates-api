#!/bin/bash

cat <<EOF > /etc/systemd/system/hotelrateapi.service
[Unit]
Description= Hotel Rates API
After=network.target

[Service]
Type=simple
WorkingDirectory=/
ExecStart=$(dirname "$0")/app/hotel-rates-api
Restart=always
RestartSec=3
Environment="ENV_VARIABLE=value"

[Install]
WantedBy=multi-user.target
EOF 
systemctl daemon-reload
systemctl start hotelrateapi.service