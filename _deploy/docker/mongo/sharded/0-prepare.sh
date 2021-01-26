#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


../network-setup.sh


# --- configsvr ---

set_name=${CONFIGSVR_REPLICA_SET_NAME}

BEGIN=${CONFIGSVR_NODE_BEGIN}
END=${CONFIGSVR_NODE_END}

node_name="${set_name}-${BEGIN}"

if [ -d "./build/${node_name}" ]; then
  exit 0
fi


for node in `seq $BEGIN $END`; do
  node_name="${set_name}-${node}"
  mkdir -p ./build/${node_name} \
  && chmod a+w ./build/${node_name}
done


# --- shardsvr ---

NUM_BEGIN=${SHARD_NUM_BEGIN}
NUM_END=${SHARD_NUM_END}

set_name_prefix=${SHARD_REPLICA_SET_NAME_PREFIX}

set_name="${set_name_prefix}-${NUM_BEGIN}"
node_name="${set_name}-${NODE_BEGIN}"

if [ -d "./build/${node_name}" ]; then
  exit 0
fi


for num in `seq $NUM_BEGIN $NUM_END`; do
  set_name="${set_name_prefix}-${num}"

  for node in `seq $NODE_BEGIN $NODE_END`; do
    node_name="${set_name}-${node}"
    mkdir -p ./build/${node_name} \
    && chmod a+w ./build/${node_name}
  done
done
