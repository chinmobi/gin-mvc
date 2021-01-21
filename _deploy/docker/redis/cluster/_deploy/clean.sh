#!/usr/bin/env bash

cd $(dirname $0) && source ../NODE.variables


if docker ps | grep -q "redis-${PORT_BEGIN}"; then
  ../stop.sh
fi


if docker ps -a | grep -q "redis-${PORT_BEGIN}"; then
  for port in `seq $PORT_BEGIN $PORT_END`; do \
    docker container rm -f redis-${port}; \
  done
fi


if [ ! -d "./${PORT_BEGIN}" ]; then
  exit 0
fi

for port in `seq $PORT_BEGIN $PORT_END`; do \
  rm -rf ./${port}; \
done
