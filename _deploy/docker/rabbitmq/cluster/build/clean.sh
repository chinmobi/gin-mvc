#!/usr/bin/env bash

cd $(dirname $0) && source ../NODE.variables


if docker ps | grep -q "rabbit-${NODE_BEGIN}"; then
  ../stop.sh
fi


if docker ps -a | grep -q "rabbit-${NODE_BEGIN}"; then
  for node in `seq $NODE_BEGIN $NODE_END`; do \
    docker container rm -f rabbit-${node}; \
  done
fi

if docker ps -a | grep -q "${RABBITMQ_HAPROXY}"; then
  docker container rm -f ${RABBITMQ_HAPROXY}
fi


if [ ! -d "./${NODE_BEGIN}" ]; then
  exit 0
fi

for node in `seq $NODE_BEGIN $NODE_END`; do \
  sudo rm -rf ./${node}; \
done

sudo rm -rf ./haproxy
