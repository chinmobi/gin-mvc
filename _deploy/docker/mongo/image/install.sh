#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


if ! docker images | grep -q "${IMAGE_REPOSITORY}"; then
  docker pull ${IMAGE_FULL_NAME}
fi
