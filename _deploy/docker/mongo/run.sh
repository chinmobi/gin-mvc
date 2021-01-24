#!/usr/bin/env bash

cd $(dirname $0) && source ./.env && source ./DB.variables


if ! docker ps | grep -q "${CONTAINER_NAME}"; then
  if docker ps -a | grep -q "${CONTAINER_NAME}"; then
    docker start ${CONTAINER_NAME}
  else
    ./volume/setup.sh
    ./network-setup.sh

    docker run -d \
      --name ${CONTAINER_NAME} \
      --network=${BACKEND_NETWORK_NAME} --hostname=${CONTAINER_NAME} \
      -p ${HOST_PORT}:${PORT} \
      -v ${VOLUME_NAME}:/bitnami/mongodb \
      -e MONGODB_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
      -e MONGODB_USERNAME=${MONGODB_USERNAME} \
      -e MONGODB_PASSWORD=${MONGODB_PASSWORD} \
      -e MONGODB_DATABASE=${MONGODB_DATABASE} \
      ${IMAGE_FULL_NAME}
  fi
fi
