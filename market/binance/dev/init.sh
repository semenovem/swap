#!/bin/bash

BIN=$(dirname $([[ $0 = /* ]] && echo "$0" || echo "$PWD/${0#./}"))
ENV=/.env123__321

if [ -f "${ENV}" ]; then
  exit 0
fi


set -o allexport
source "${BIN}/config.env"
set +o allexport


# ставим признак того, что инициализация была
echo "initialized" > "${ENV}"
