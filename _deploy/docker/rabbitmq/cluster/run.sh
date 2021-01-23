#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if docker ps | grep -q "rabbit-${NODE_BEGIN}"; then
  exit 0
fi

if ! docker ps -a | grep -q "rabbit-${NODE_BEGIN}"; then
  exit 1
fi


for node in `seq $NODE_END -1 $NODE_BEGIN`; do \
  docker start rabbit-${node}; \
done

docker start ${RABBITMQ_HAPROXY}
