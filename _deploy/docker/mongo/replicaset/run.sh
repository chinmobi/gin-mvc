#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


set_name=${REPLICA_SET_NAME}

node_name="${set_name}-${NODE_BEGIN}"

if docker ps | grep -q "${node_name}"; then
  exit 0
fi

if ! docker ps -a | grep -q "${node_name}"; then
  exit 1
fi


for node in `seq $NODE_BEGIN $NODE_END`; do \
  node_name="${set_name}-${node}"; \
  docker start ${node_name}; \
done


node_name="${set_name}-arbiter"

if docker ps -a | grep -q "${node_name}"; then
  docker start "${node_name}"
fi
