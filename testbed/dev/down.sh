#!/bin/bash

echo "############################################################"
echo "# Остановка тестового стенда bot dev                       #"
echo "############################################################"

BIN=$(dirname $([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}"))
source "${BIN}/util.sh"

docker stack rm \
  "${__STACK_PREFIX__}-postgres"
