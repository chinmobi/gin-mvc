#!/usr/bin/env bash

cd $(dirname $0) && source ./.env && source ./DB.variables


if docker ps | grep -q "${CONTAINER_NAME}"; then
  #docker exec -it ${CONTAINER_NAME} bash
  docker run -it --rm --network=${BACKEND_NETWORK_NAME} \
  -v ${HOME_MOUNT}:/home/mongodb \
  ${IMAGE_FULL_NAME} \
  mongo --host ${CONTAINER_NAME} \
    -u ${MONGO_INITDB_ROOT_USERNAME} \
    -p ${MONGO_INITDB_ROOT_PASSWORD} \
    --authenticationDatabase admin \
    admin
elif docker images | grep -q "${IMAGE_REPOSITORY}"; then
  docker run -it --rm ${IMAGE_FULL_NAME} bash
fi
