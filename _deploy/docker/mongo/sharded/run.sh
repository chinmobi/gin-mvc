#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../.env


# --- configsvr ---

set_name=${CONFIGSVR_REPLICA_SET_NAME}

BEGIN=${CONFIGSVR_NODE_BEGIN}
END=${CONFIGSVR_NODE_END}

node_name="${set_name}-${BEGIN}"

if ! docker ps | grep -q "${node_name}"; then
  if docker ps -a | grep -q "${node_name}"; then
    for node in `seq $BEGIN $END`; do \
      node_name="${set_name}-${node}"; \
      docker start ${node_name}; \
    done
  fi
fi


# --- mongos ---

node_name=${CONTAINER_NAME}

if ! docker ps | grep -q "${node_name}"; then
  if docker ps -a | grep -q "${node_name}"; then
    docker start "${node_name}"
  fi
fi


# --- shardsvr ---

NUM_BEGIN=${SHARD_NUM_BEGIN}
NUM_END=${SHARD_NUM_END}

set_name_prefix=${SHARD_REPLICA_SET_NAME_PREFIX}

set_name="${set_name_prefix}-${NUM_BEGIN}"
node_name="${set_name}-${NODE_BEGIN}"

if ! docker ps | grep -q "${node_name}"; then
  if docker ps -a | grep -q "${node_name}"; then
    for num in `seq $NUM_BEGIN $NUM_END`; do
      set_name="${set_name_prefix}-${num}"

      for node in `seq $NODE_BEGIN $NODE_END`; do \
        node_name="${set_name}-${node}"; \
        docker start "${node_name}"; \
      done

      node_name="${set_name}-arbiter"
      docker start "${node_name}"
    done
  fi
fi
