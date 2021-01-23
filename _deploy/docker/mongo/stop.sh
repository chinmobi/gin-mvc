#!/usr/bin/env bash

cd $(dirname $0) && source ./.env


if docker ps | grep -q "${CONTAINER_NAME}"; then
  docker stop ${CONTAINER_NAME}
fi

if docker ps -a | grep -q "${CONTAINER_NAME}"; then
  docker container rm -f ${CONTAINER_NAME} && \
  docker container prune -f
fi
