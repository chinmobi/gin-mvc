#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


set_name=${REPLICA_SET_NAME}

node_name="${set_name}-arbiter"

if docker ps | grep -q "${node_name}"; then
  docker stop "${node_name}"
fi


node_name="${set_name}-${NODE_BEGIN}"

if ! docker ps | grep -q "${node_name}"; then
  exit 0
fi

for node in `seq $NODE_END -1 $NODE_BEGIN`; do \
  node_name="${set_name}-${node}"; \
  docker stop ${node_name}; \
done
