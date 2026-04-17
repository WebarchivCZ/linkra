#!/bin/bash

function usage() {
	cat <<-EOF
		usage: run-env.sh [ start | upgrade ] env_name

		If first argument is 'start' then this script runs docker compose up
		for specified env_name.
		env_name is name of directory in the ./docker directory.
		It will run it as project "linkraENV_NAME".

		If first argument is 'upgrade' it will rebuild, recreate and restart
		the docker compose project.
	EOF
}

if [ $# -eq 2 ]; then
	COMMAND="$1"
	ENV_NAME="$2"
	ENV_DIR="./docker/$ENV_NAME"
	PROJECT="linkra-$ENV_NAME"
	if [ $COMMAND = "start" ]; then
	# echo "Using dir $ENV_DIR and project $PROJECT"
		docker compose -f "$ENV_DIR/docker-compose.yml" -p "$PROJECT" up -d
	elif [ $COMMAND = "upgrade" ]; then
	# echo "Using dir $ENV_DIR and project $PROJECT"
		docker compose -f "$ENV_DIR/docker-compose.yml" -p "$PROJECT" up -d --build --force-recreate
	fi
else
	echo "Wrong number of arguments!"
	usage
	exit 1
fi