#!/usr/bin/env bash

cd $(dirname $0) && source ./.env


if ! docker ps -a | grep -q "${CONTAINER_NAME}"; then
  ./volume/setup.sh
  ./network-setup.sh

  docker-compose up -d
fi
