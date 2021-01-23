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
      -v ${VOLUME_NAME}:/data/db \
      -v ${CONFIG_VOLUME_NAME}:/data/configdb \
      -e MONGO_INITDB_ROOT_USERNAME=${MONGO_INITDB_ROOT_USERNAME} \
      -e MONGO_INITDB_ROOT_PASSWORD=${MONGO_INITDB_ROOT_PASSWORD} \
      ${IMAGE_FULL_NAME}
  fi
fi
