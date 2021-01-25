#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../DB.variables \
&& source ../.env


set_name=${REPLICA_SET_NAME}

if [ "$#" -gt 0 ]; then
  set_name=$1
fi

node_name="${set_name}-${NODE_BEGIN}"


if docker ps -a | grep -q "${node_name}"; then
  exit 0
fi


./0-prepare.sh


docker run -d \
  --name ${node_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${node_name} \
  -v $(pwd)/build/${node_name}:/bitnami/mongodb \
  -e MONGODB_REPLICA_SET_MODE=primary \
  -e MONGODB_REPLICA_SET_NAME=${set_name} \
  -e MONGODB_PORT_NUMBER=${PORT} \
  -e MONGODB_ADVERTISED_HOSTNAME=${node_name} \
  -e MONGODB_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${IMAGE_FULL_NAME}
