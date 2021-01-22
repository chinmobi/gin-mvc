#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if ! docker ps | grep -q "rabbit-${PORT_END}"; then
  exit 0
fi


docker stop ${RABBITMQ_HAPROXY}

for port in `seq $PORT_BEGIN $PORT_END`; do \
  docker stop rabbit-${port}; \
done
