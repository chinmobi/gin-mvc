#!/usr/bin/env bash

cd $(dirname $0) && source ./.env


BUILD_OPT="--build"
if docker images | grep -q "${IMAGE_REPOSITORY}"; then
  BUILD_OPT=""
fi

if ! docker ps -a | grep -q "${CONTAINER_NAME}"; then
  ./networks/setup.sh

  docker-compose up ${BUILD_OPT} -d
fi
