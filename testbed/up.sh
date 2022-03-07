#!/bin/bash

echo "############################################################"
echo "# Testbed launch                                           #"
echo "############################################################"

BIN=$(dirname "$([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}")")
source "${BIN}/util.sh"


NET=$1

choiceTestbed "$NET"
res=$?
[ $res -eq 0 ] && echo "ERR: did not define the name of the docker network" && exit 1
NAME_NET=$(getNetByNum $res)

echo "docker network name:  $NAME_NET"


#docker stack deploy \
#  -c "${BIN}/docker-postgres.yaml" "${__STACK_PREFIX__}-postgres"
