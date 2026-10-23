#!/bin/bash

function usage() {
	cat <<-EOF
		usage: create-new-env.sh env_name [file with variables]

		This script will create new directory containing docker-compose file
		and worker configuration json file necessary to start a linkra application
		using docker compose. The directory created will be "./docker/env_name"

		Set values for variables directly in this script or use external file.

		If external file is used, it should be a bash script containing variable
		declarations in usual format (NAME="value"). Create the file then pass
		the filename as argument for this script.
	EOF
}

if [ "$1" = "-h" ] || [ "$1" = "--help" ]; then
	usage
	exit 0
fi

if [ $# -eq 0 ]; then
	echo "Not enough arguments!"
	echo
	usage
	exit 1
elif [ $# -eq 1 ]; then
# -------------------------- Change this ---------------------------------------
# If you want an external config then copy this part and paste it into new file.
# Or just edit this file directly.
# Note that this settings are docker specific and don't completely mirror
# env variables used for running as native process.

# --- Server configuration ---
# Where to mount the DB file (relative to docker-compose.yaml)
SQLITE_DB_MOUNT="../../storage.db"
# Where will be the app accessed (public URL)
SERVER_HOST="http://localhost:8080"
# Port
SERVER_PORT="8080"

HOST_UID="1000"
HOST_GID="1000"

# --- Worker configuration ---
# Where to store captures (relative to docker-compose.yaml)
CAPTURES_MOUNT="../../captures/"
# Do we want to discard captures
DISCARD_ARCHIVE_FILES="false"
# Settings for Scoop
CAPTURE_SETTINGS='{
    "chromiumSandbox": false
  }'

# ------------------------------------------------------------------------------
elif [ $# -eq 2 ]; then
	VAR_FILE="$2"
	if [ -f "$VAR_FILE" ]; then
		. "$VAR_FILE"
	else
		echo "The file $VAR_FILE does not exist!"
		echo
		usage
		exit 1
	fi
else
	echo "Too many arguments!"
	echo
	usage
	exit 1
fi

ENV_DIRECTORY="./docker/$1"
mkdir -p "$ENV_DIRECTORY"

# Create docker compose
cat > "$ENV_DIRECTORY/docker-compose.yml" <<EOF
networks:
  queue-network: {}

services:
  linkra-server:
    build:
      context: ../..
      args:
        TARGET_UID: ${HOST_UID}
        TARGET_GID: ${HOST_GID}
    ports:
      - $SERVER_PORT:8080
    volumes:
      - $SQLITE_DB_MOUNT:/mnt/storage.db:rw
    depends_on:
      - valkey
    networks:
      queue-network: {}
    environment:
      - DB_PATH=/mnt/storage.db
      - VALKEY_ADDR=queue
      - VALKEY_PORT=6379
      - SERVER_ADDRESS=0.0.0.0:8080
      - SERVER_HOST=$SERVER_HOST
    restart: on-failure:3

  linkra-worker:
    build:
      context: ../../workers/scoop-worker
    volumes:
      - ./worker-config.json:/app/config.json:ro
      - $CAPTURES_MOUNT:/app/captures/:rw
    networks:
      queue-network: {}
    depends_on:
      - valkey
    restart: on-failure:3
    cap_drop:
      - ALL
    security_opt:
      - no-new-privileges:true

  valkey:
    image: valkey/valkey:8.0-alpine
    command: >
      valkey-server
      --save 60 1
    networks:
      queue-network:
        aliases:
          - queue
EOF

cat > "$ENV_DIRECTORY/worker-config.json" <<EOF
{
  "discardArchiveFiles": $DISCARD_ARCHIVE_FILES,
  "outputDir": "/app/captures/",
  "valkeyUrl": "redis://queue:6379",
  "captureSettings": $CAPTURE_SETTINGS
}
EOF
