#!/bin/bash

echo "############################################################"
echo "# Запуск / остановка стека                                 #"
echo "############################################################"

BIN=$(dirname $([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}"))
source "${BIN}/util.sh"

NAME=

for p in "$@"; do
  [ "$p" == 'stop' ] && OPER="stop" && continue
  [ "$p" == '-' ] && OPER="stop" && continue
  [ "$p" == 'down' ] && OPER="stop" && continue
  [ "$p" == 'restart' ] && OPER="restart" && continue

  [ -f "${BIN}/docker-${p}.yaml" ] && NAME=$p
done

[ -z $OPER ] && OPER="start"

STACK="${__STACK_PREFIX__}-${NAME}"
FILE="${BIN}/docker-${NAME}.yaml"

[ ! -f "$FILE" ] && echo "Error: docker file does not exist: $FILE" && exit 1

USE=

if [ $OPER == "start" ]; then
  USE=1
  docker stack deploy \
    -c "$FILE" \
    "$STACK"
fi

if [ $OPER == "stop" ]; then
  USE=1
  docker stack rm "$STACK"
fi

if [ $OPER == "restart" ]; then
  USE=1
  bash "$0" "stop" "$NAME"
  sleep 1
  bash "$0" "start" "$NAME"
fi

[ -z "$USE" ] && echo "Error: unknown command passed: $OPER"
