#!/usr/bin/env bash

cd $(dirname $0) && source ./.env


if docker ps | grep -q "${CONTAINER_NAME}"; then
  docker exec -it ${CONTAINER_NAME} \
  psql -h ${CONTAINER_NAME} -U postgres --password
elif docker images | grep -q "${IMAGE_REPOSITORY}"; then
  docker run -it --rm ${IMAGE_FULL_NAME} bash
fi
