#!/bin/bash

set -e

LOG_DIR=/var/log/httpd
VOLUME=/home/medora/log/httpd
STATUS="200,201,202,203,204,300,301,302"

OPTIONS="-f apache_combined --overwrite -t log -o ${LOG_DIR}/access_log"

mkdir -p $VOLUME

sudo docker run \
	-d \
	-e STATUS_LIMIT=$STATUS \
	-e MAX_SLEEP="60" \
	--restart=unless-stopped \
	-v $VOLUME:$LOG_DIR \
	bmedora/flog $OPTIONS
