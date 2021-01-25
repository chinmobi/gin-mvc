#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


if ! docker images | grep -q "${IMAGE_REPOSITORY}"; then
  docker pull ${IMAGE_FULL_NAME}
fi

if ! docker images | grep -q "${SHARDED_IMAGE_REPOSITORY}"; then
  docker pull ${SHARDED_IMAGE_FULL_NAME}
fi
