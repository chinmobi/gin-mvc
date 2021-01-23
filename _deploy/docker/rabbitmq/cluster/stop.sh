#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if ! docker ps | grep -q "rabbit-${NODE_END}"; then
  exit 0
fi


docker stop ${RABBITMQ_HAPROXY}

for node in `seq $NODE_BEGIN $NODE_END`; do \
  docker stop rabbit-${node}; \
done
