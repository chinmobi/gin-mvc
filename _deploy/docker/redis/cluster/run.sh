#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if docker ps | grep -q "redis-${PORT_BEGIN}"; then
  exit 0
fi

if ! docker ps -a | grep -q "redis-${PORT_BEGIN}"; then
  exit 1
fi


for port in `seq $PORT_BEGIN $PORT_END`; do \
  docker start redis-${port}; \
done
