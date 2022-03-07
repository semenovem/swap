#!/bin/bash

BIN=$(dirname "$([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}")")
source "${BIN}/util.sh"

NET=$1

choiceNet "$NET"
res=$?
[ $res -eq 0 ] && echo "ERR: did not define the name of the docker network" && exit 1
NAME_NET=$(getNetByNum $res)

echo "docker network name:  $NAME_NET"

[ "$(docker network ls -f name="$NAME_NET" -q)" ] && exit 0

docker network create \
  --driver overlay \
  --attachable \
  "$NAME_NET"
