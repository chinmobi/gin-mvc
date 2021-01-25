#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


set_name=${REPLICA_SET_NAME}

if [ "$#" -gt 0 ]; then
  set_name=$1
fi


../network-setup.sh


node_name="${set_name}-${NODE_BEGIN}"

if [ -d "./build/${node_name}" ]; then
  exit 0
fi

for node in `seq $NODE_BEGIN $NODE_END`; do \
  node_name="${set_name}-${node}"; \
  mkdir -p ./build/${node_name} \
  && chmod a+w ./build/${node_name}; \
done
