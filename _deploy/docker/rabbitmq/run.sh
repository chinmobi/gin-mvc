#!/usr/bin/env bash

cd $(dirname $0) && source ./.env


if ! docker ps | grep -q "${CONTAINER_NAME}"; then
  if docker ps -a | grep -q "${CONTAINER_NAME}"; then
    docker start ${CONTAINER_NAME}
  else
    ./volume/setup.sh
    ./network-setup.sh

    docker run -d \
      --name ${CONTAINER_NAME} \
      --network=${BACKEND_NETWORK_NAME} --hostname=${CONTAINER_NAME} \
      -p ${HOST_PORT}:${PORT} -p ${MGMT_HOST_PORT}:${MGMT_PORT} \
      -v ${VOLUME_NAME}:/var/lib/rabbitmq \
      ${MANAGEMENT_IMAGE_FULL_NAME}
  fi
fi
