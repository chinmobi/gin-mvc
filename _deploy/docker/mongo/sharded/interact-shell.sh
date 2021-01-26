#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../DB.variables \
&& source ../.env


node_name=${CONTAINER_NAME}

if [ "$#" -gt 0 ]; then
  node_name=$1
fi


if docker ps | grep -q "${node_name}"; then
  docker run -it --rm --network=${BACKEND_NETWORK_NAME} \
  ${IMAGE_FULL_NAME} \
  mongo --host ${node_name} \
    -u root \
    -p ${MONGODB_ROOT_PASSWORD} \
    admin
fi
