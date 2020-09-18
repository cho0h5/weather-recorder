#!/bin/bash

#get path
RELATIVE_DIR=`dirname "$0"`
cd $RELATIVE_DIR
SHELL_PATH=`pwd -P`

#build
go build .

#
DESCRIPTION="weather-recoder"
BINARY_NAME="collect-server"
echo "[Unit]
Description=$DESCRIPTION
After=network.target

[Service]
WorkingDirectory=$SHELL_PATH
ExecStart=$SHELL_PATH/$BINARY_NAME

[Install]
WantedBy=multi-user.target" >> /etc/systemd/system/$DESCRIPTION.service

#systemd
systemctl enable $DESCRIPTION
systemctl start $DESCRIPTION
IS_ACTIVE=`systemctl is-active $DESCRIPTION`
echo "$DESCRIPTION $IS_ACTIVE"
