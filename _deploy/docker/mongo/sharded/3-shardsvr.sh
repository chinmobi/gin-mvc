#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../DB.variables \
&& source ../.env


mongos_node_name=${CONTAINER_NAME}

if ! docker ps | grep -q "${mongos_node_name}"; then
  exit 1
fi


NUM_BEGIN=${SHARD_NUM_BEGIN}
NUM_END=${SHARD_NUM_END}

set_name_prefix=${SHARD_REPLICA_SET_NAME_PREFIX}

set_name="${set_name_prefix}-${NUM_BEGIN}"
node_name="${set_name}-${NODE_BEGIN}"

if docker ps -a | grep -q "${node_name}"; then
  exit 0
fi


# --- primary ---

for num in `seq $NUM_BEGIN $NUM_END`; do
  set_name="${set_name_prefix}-${num}"
  node_name="${set_name}-${NODE_BEGIN}"

  docker run -d \
  --name ${node_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${node_name} \
  -v $(pwd)/build/${node_name}:/bitnami \
  -e MONGODB_SHARDING_MODE=shardsvr \
  -e MONGODB_MONGOS_HOST=${mongos_node_name} \
  -e MONGODB_REPLICA_SET_MODE=primary \
  -e MONGODB_REPLICA_SET_NAME=${set_name} \
  -e MONGODB_PORT_NUMBER=${PORT} \
  -e MONGODB_ADVERTISED_HOSTNAME=${node_name} \
  -e MONGODB_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${SHARDED_IMAGE_FULL_NAME}
done


# --- secondaries ---

BEGIN=$(($NODE_BEGIN+1))

for num in `seq $NUM_BEGIN $NUM_END`; do
  set_name="${set_name_prefix}-${num}"
  primary_name="${set_name}-${NODE_BEGIN}"

  for node in `seq $BEGIN $NODE_END`; do \
  node_name="${set_name}-${node}"; \
  docker run -d \
  --name ${node_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${node_name} \
  -v $(pwd)/build/${node_name}:/bitnami \
  -e MONGODB_SHARDING_MODE=shardsvr \
  -e MONGODB_REPLICA_SET_MODE=secondary \
  -e MONGODB_REPLICA_SET_NAME=${set_name} \
  -e MONGODB_PORT_NUMBER=${PORT} \
  -e MONGODB_ADVERTISED_HOSTNAME=${node_name} \
  -e MONGODB_PRIMARY_HOST=${primary_name} \
  -e MONGODB_PRIMARY_PORT_NUMBER=${PORT} \
  -e MONGODB_PRIMARY_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${SHARDED_IMAGE_FULL_NAME}; \
  done
done


# --- arbiter ---

for num in `seq $NUM_BEGIN $NUM_END`; do
  set_name="${set_name_prefix}-${num}"
  primary_name="${set_name}-${NODE_BEGIN}"

  node_name="${set_name}-arbiter"

  docker run -d \
  --name ${node_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${node_name} \
  -e MONGODB_SHARDING_MODE=shardsvr \
  -e MONGODB_REPLICA_SET_MODE=arbiter \
  -e MONGODB_REPLICA_SET_NAME=${set_name} \
  -e MONGODB_ADVERTISED_HOSTNAME=${node_name} \
  -e MONGODB_PRIMARY_HOST=${primary_name} \
  -e MONGODB_PRIMARY_PORT_NUMBER=${PORT} \
  -e MONGODB_PRIMARY_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${SHARDED_IMAGE_FULL_NAME}
done
