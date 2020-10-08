#!/bin/bash

set -e

LOG_DIR=/var/log/httpd
VOLUME=/home/medora/log/httpd
STATUS="400,404,500,503"

OPTIONS="-f apache_combined --overwrite -t log -o ${LOG_DIR}/access_log.bad"

mkdir -p $VOLUME

sudo docker run \
	-d \
	-e STATUS_LIMIT=$STATUS \
	-e MAX_SLEEP="120" \
	--restart=unless-stopped \
	-v $VOLUME:$LOG_DIR \
	bmedora/flog $OPTIONS
