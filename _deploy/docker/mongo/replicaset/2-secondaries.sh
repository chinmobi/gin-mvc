#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../DB.variables \
&& source ../.env


set_name=${REPLICA_SET_NAME}

if [ "$#" -gt 0 ]; then
  set_name=$1
fi


primary_name="${set_name}-${NODE_BEGIN}"


if ! docker ps | grep -q "${primary_name}"; then
  exit 1
fi


BEGIN=$(($NODE_BEGIN+1))

node_name="${set_name}-${BEGIN}"

if docker ps | grep -q "${node_name}"; then
  exit 0
fi

for node in `seq $BEGIN $NODE_END`; do \
  node_name="${set_name}-${node}"; \
  docker run -d \
  --name ${node_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${node_name} \
  -v $(pwd)/build/${node_name}:/bitnami/mongodb \
  -e MONGODB_REPLICA_SET_MODE=secondary \
  -e MONGODB_REPLICA_SET_NAME=${set_name} \
  -e MONGODB_PORT_NUMBER=${PORT} \
  -e MONGODB_ADVERTISED_HOSTNAME=${node_name} \
  -e MONGODB_INITIAL_PRIMARY_HOST=${primary_name} \
  -e MONGODB_INITIAL_PRIMARY_PORT_NUMBER=${PORT} \
  -e MONGODB_INITIAL_PRIMARY_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${IMAGE_FULL_NAME}; \
done
