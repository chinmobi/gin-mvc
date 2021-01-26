#!/usr/bin/env bash

cd $(dirname $0) && source ../NODE.variables \
&& source ../../.env


set_name=${CONFIGSVR_REPLICA_SET_NAME}
BEGIN=${CONFIGSVR_NODE_BEGIN}
node_name="${set_name}-${BEGIN}"

if docker ps | grep -q "${node_name}"; then
  ../stop.sh
fi


# --- shardsvr ---

NUM_BEGIN=${SHARD_NUM_BEGIN}
NUM_END=${SHARD_NUM_END}

set_name_prefix=${SHARD_REPLICA_SET_NAME_PREFIX}

set_name="${set_name_prefix}-${NUM_BEGIN}"
node_name="${set_name}-${NODE_BEGIN}"

if docker ps -a | grep -q "${node_name}"; then
  for num in `seq $NUM_END -1 $NUM_BEGIN`; do
    set_name="${set_name_prefix}-${num}"

    node_name="${set_name}-arbiter"
    docker container rm -f "${node_name}"

    for node in `seq $NODE_END -1 $NODE_BEGIN`; do
      node_name="${set_name}-${node}"
      docker container rm -f "${node_name}"
    done
  done
fi


set_name="${set_name_prefix}-${NUM_BEGIN}"
node_name="${set_name}-${NODE_BEGIN}"

if [ -d "./${node_name}" ]; then
  for num in `seq $NUM_BEGIN $NUM_END`; do
    set_name="${set_name_prefix}-${num}"

    for node in `seq $NODE_BEGIN $NODE_END`; do
      node_name="${set_name}-${node}"
      sudo rm -rf ./${node_name}
    done
  done
fi


# --- mongos ---

node_name=${CONTAINER_NAME}

if docker ps -a | grep -q "${node_name}"; then
  docker container rm -f "${node_name}"
fi


# --- configsvr ---

set_name=${CONFIGSVR_REPLICA_SET_NAME}

BEGIN=${CONFIGSVR_NODE_BEGIN}
END=${CONFIGSVR_NODE_END}

node_name="${set_name}-${BEGIN}"

if docker ps -a | grep -q "${node_name}"; then
  for node in `seq $END -1 $BEGIN`; do
    node_name="${set_name}-${node}"
    docker container rm -f ${node_name}
  done
fi


node_name="${set_name}-${BEGIN}"

if [ -d "./${node_name}" ]; then
  for node in `seq $BEGIN $END`; do
    node_name="${set_name}-${node}"
    sudo rm -rf ./${node_name}
  done
fi
