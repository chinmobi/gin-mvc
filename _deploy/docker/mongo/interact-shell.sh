#!/usr/bin/env bash

cd $(dirname $0) && source ./.env && source ./DB.variables


if docker ps | grep -q "${CONTAINER_NAME}"; then
  #docker exec -it ${CONTAINER_NAME} bash
  docker run -it --rm --network=${BACKEND_NETWORK_NAME} \
  ${IMAGE_FULL_NAME} \
  mongo --host ${CONTAINER_NAME} \
    -u ${MONGODB_USERNAME} \
    -p ${MONGODB_PASSWORD} \
    ${MONGODB_DATABASE}
elif docker images | grep -q "${IMAGE_REPOSITORY}"; then
  docker run -it --rm ${IMAGE_FULL_NAME} bash
fi
