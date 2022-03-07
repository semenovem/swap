#!/bin/bash

BIN=$(dirname $([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}"))
bash "${BIN}/init.sh"

go get


while true; do
  echo
  echo "###########################################################"
  echo "# Старт приложения в DEV MODE                             #"
  echo "# [ctrl + c] - два раза подряд для выхода                 #"
  echo "###########################################################"

  set -o allexport
  source "${BIN}/config.env"
  source "${BIN}/../project.properties"
  set +o allexport

  export ENV_PKG_APP_BUILD_DEV_MODE_VERSION=${VERSION}
  export ENV_PKG_APP_BUILD_DEV_MODE_HASH_COMMIT=$(git rev-parse --verify HEAD)
# версия git в образе - < 2.22
# если будет выше - используй `git branch --show-current`
  export ENV_PKG_APP_BUILD_DEV_MODE_BRANCH=$(git rev-parse --abbrev-ref HEAD)
  export ENV_PKG_APP_BUILD_DEV_MODE_TIME=$(date)

  go run *.go

  sleep 1
done
