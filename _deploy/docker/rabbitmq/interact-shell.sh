#!/usr/bin/env bash

cd $(dirname $0) && source ./.env


if docker ps | grep -q "${CONTAINER_NAME}"; then
  docker exec -it ${CONTAINER_NAME} /bin/bash
elif docker images | grep -q "${IMAGE_REPOSITORY}"; then
  docker run -it --rm ${MANAGEMENT_IMAGE_FULL_NAME} /bin/bash
fi
