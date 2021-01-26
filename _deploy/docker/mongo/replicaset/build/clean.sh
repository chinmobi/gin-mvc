#!/usr/bin/env bash

cd $(dirname $0) && source ../NODE.variables


set_name=${REPLICA_SET_NAME}

node_name="${set_name}-${NODE_BEGIN}"

if docker ps | grep -q "${node_name}"; then
  ../stop.sh
fi


if docker ps -a | grep -q "${node_name}"; then
  for node in `seq $NODE_BEGIN $NODE_END`; do \
    node_name="${set_name}-${node}"; \
    docker container rm -f ${node_name}; \
  done
fi

node_name="${set_name}-arbiter"

if docker ps -a | grep -q "${node_name}"; then
  docker container rm -f ${node_name}
fi


node_name="${set_name}-${NODE_BEGIN}"

if [ ! -d "./${node_name}" ]; then
  exit 0
fi

for node in `seq $NODE_BEGIN $NODE_END`; do \
  node_name="${set_name}-${node}"; \
  sudo rm -rf ./${node_name}; \
done
