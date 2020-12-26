#!/bin/sh
set -eo pipefail

source ./.env


if [ "${1:0:1}" = '-' ]; then
  set -- start-app "$@"
fi

if [ "$1" = 'start-app' ]; then
  exec ./bin/${APP_NAME}
else
  exec "$@"
fi
