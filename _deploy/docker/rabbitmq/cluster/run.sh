#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if docker ps | grep -q "rabbit-${PORT_BEGIN}"; then
  exit 0
fi

if ! docker ps -a | grep -q "rabbit-${PORT_BEGIN}"; then
  exit 1
fi


for port in `seq $PORT_END -1 $PORT_BEGIN`; do \
  docker start rabbit-${port}; \
done

docker start ${RABBITMQ_HAPROXY}
