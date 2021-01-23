#!/usr/bin/env bash

cd $(dirname $0) \
&& source ../NODE.variables \
&& source ../.env


if docker ps | grep -q "pg-${NODE_BEGIN}"; then
  ../stop.sh
fi


if docker ps -a | grep -q "pg-${NODE_BEGIN}"; then
  for node in `seq $NODE_BEGIN $NODE_END`; do \
    docker container rm -f pg-${node}; \
  done
fi

if docker ps -a | grep -q "${PGPOOL_CONTAINER_NAME}"; then
  docker container rm -f ${PGPOOL_CONTAINER_NAME}
fi


if [ ! -d "./pg-${NODE_BEGIN}" ]; then
  exit 0
fi

for node in `seq $NODE_BEGIN $NODE_END`; do \
  sudo rm -rf ./pg-${node}; \
done
