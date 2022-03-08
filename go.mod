module ydgo

go 1.17

[Unit]
Description=goweb

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/user/go/go-web/main

[Install]
WantedBy=multi-user.target