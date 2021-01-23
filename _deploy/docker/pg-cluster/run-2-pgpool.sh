#!/usr/bin/env bash

cd $(dirname $0) \
&& source ./NODE.variables \
&& source ./.env


if docker ps | grep -q "${PGPOOL_CONTAINER_NAME}"; then
  exit 0
fi

if ! docker ps -a | grep -q "${PGPOOL_CONTAINER_NAME}"; then
  exit 1
fi

if ! docker ps | grep -q "pg-${NODE_END}"; then
  exit 1
fi


docker start ${PGPOOL_CONTAINER_NAME}
