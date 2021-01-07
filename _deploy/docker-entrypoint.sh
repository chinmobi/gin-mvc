#!/bin/sh
set -eo pipefail

source ./app.env


if [ "${1:0:1}" = '-' ]; then
  set -- start-app "$@"
fi

if [ "$1" = 'start-app' ]; then
  exec ./bin/${APP_BIN_NAME}
else
  exec "$@"
fi
