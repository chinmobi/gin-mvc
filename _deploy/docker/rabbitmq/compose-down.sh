#!/usr/bin/env bash

cd $(dirname $0) && source ./.env


if docker ps | grep -q "${CONTAINER_NAME}"; then
  docker-compose down -v
fi
