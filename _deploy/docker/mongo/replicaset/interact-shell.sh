#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../DB.variables \
&& source ../.env


set_name=${REPLICA_SET_NAME}
node=${NODE_BEGIN}

if [ "$#" -gt 0 ]; then
  node=$1
fi

node_name="${set_name}-${node}"


if docker ps | grep -q "${node_name}"; then
  docker run -it --rm --network=${BACKEND_NETWORK_NAME} \
  ${IMAGE_FULL_NAME} \
  mongo --host ${node_name} \
    -u root \
    -p ${MONGODB_ROOT_PASSWORD} \
    admin
fi
