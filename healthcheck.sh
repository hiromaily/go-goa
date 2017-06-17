#!/bin/bash

echo "[HealthCheck] Run healthcheck"

DOCKER_API_VER=v1.30
HEALTH_API=http://webserver/api/_ah/health
CONTAINER_NAME=gogoa_webserver_1

# call health check API
RESULT=`curl -s -o /dev/null "${HEALTH_API}" -w  '{"http_code" : "%{http_code}","content_type" : "%{content_type}"}' | jq '.content_type'`
if echo $RESULT | grep -q application/json; then
	echo '[HealthCheck] It works!'
elif echo $RESULT | grep -q html; then
	echo '[HealthCheck] Not working'
	# Check ID of server container
	SERVER_PID=$(curl -s --unix-socket /var/run/docker.sock http:/${DOCKER_API_VER}/containers/json?filters=%7b%22name%22%3a%20%5b%22${CONTAINER_NAME}%22%5d%7d | jq -r '.[0] | .Id')
	EXIT_STATUS=$?
	if [ $EXIT_STATUS -gt 0 ]; then
		exit 0
	fi

	# Check server status whether restart is ongoing or not
	SERVER_STATE=$(curl -s --unix-socket /var/run/docker.sock http:/${DOCKER_API_VER}/containers/json?filters=%7b%22name%22%3a%20%5b%22${CONTAINER_NAME}%22%5d%7d | jq -r '.[0] | .State')
	EXIT_STATUS=$?
	if [ $EXIT_STATUS -gt 0 ]; then
		exit 0
	fi

	#SERVER_STATE may be `running` or `restarting` or `exited` or `dead` or `paused`
	if !(echo $SERVER_STATE | grep -q restarting); then
		# Restart
		curl --unix-socket /var/run/docker.sock -X POST http:/${DOCKER_API_VER}/containers/${SERVER_PID}/restart
	fi
else
	echo '[HealthCheck] Skip'
fi
