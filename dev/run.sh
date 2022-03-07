#!/bin/bash

BIN=$(dirname "$([[ $0 == /* ]] && echo "$0" || echo "$PWD/${0#./}")")

go get

while true; do
  bash "${BIN}/run2.sh"
  RET=$?

#  echo "exit code: $RET"

  [ "$RET" -eq 2 ] && sleep 5

  [ "$RET" -eq 1 ] && sleep 2

#  [ "$RET" -eq 0 ] && break
done
