#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../DB.variables \
&& source ../.env


cfg_set_name=${CONFIGSVR_REPLICA_SET_NAME}

BEGIN=${CONFIGSVR_NODE_BEGIN}

cfg_primary_name="${cfg_set_name}-${BEGIN}"


if ! docker ps | grep -q "${cfg_primary_name}"; then
  exit 1
fi


node_name=${CONTAINER_NAME}

if docker ps -a | grep -q "${node_name}"; then
  exit 0
fi


docker run -d \
  --name ${node_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${node_name} \
  -p ${HOST_PORT}:${PORT} \
  -e MONGODB_SHARDING_MODE=mongos \
  -e MONGODB_ADVERTISED_HOSTNAME=${node_name} \
  -e MONGODB_CFG_PRIMARY_HOST=${cfg_primary_name} \
  -e MONGODB_CFG_REPLICA_SET_NAME=${cfg_set_name} \
  -e MONGODB_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${SHARDED_IMAGE_FULL_NAME}
