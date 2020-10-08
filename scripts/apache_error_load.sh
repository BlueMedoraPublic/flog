#!/bin/bash

set -e

LOG_DIR=/var/log/httpd
VOLUME=/home/medora/log/httpd

OPTIONS="-f apache_error --overwrite -t log -o ${LOG_DIR}/error_log --loop -s 1 -d 5s -n 5"

mkdir -p $VOLUME

sudo docker run \
	-d \
	-e ERROR_LOGS="enable" \
	-e MAX_SLEEP="500" --restart=unless-stopped \
	-v $VOLUME:$LOG_DIR \
	bmedora/flog $OPTIONS
