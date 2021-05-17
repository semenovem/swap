#!/bin/bash

echo "############################################################"
echo "# Запуск тестового стенда bot dev                          #"
echo "############################################################"

BIN=$(dirname $([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}"))
source "${BIN}/util.sh"

docker stack deploy \
  -c "${BIN}/docker-postgres.yaml" "${__STACK_PREFIX__}-postgres"
