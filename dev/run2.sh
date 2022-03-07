#!/bin/bash

BIN=$(dirname "$([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}")")

echo
echo "###########################################################"
echo "# Start the application in DEV MODE                       #"
echo "# [ctrl + c] - press to exit                              #"
echo "# [q]        - press for rerun                            #"
echo "###########################################################"

set -o allexport
source "${BIN}/../app/project.properties"
set +o allexport

export ENV_FILE_CONFIG=/repo/dev/cfg/config.yaml

go run -ldflags="-X 'main.buildVersion=${APP_VER}'" *.go
exit $?
