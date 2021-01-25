#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


if docker images | grep -q "${IMAGE_REPOSITORY}"; then
  docker image rm -f ${IMAGE_FULL_NAME}
fi
