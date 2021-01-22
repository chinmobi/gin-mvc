#!/usr/bin/env bash

cd $(dirname $0) && source ../NODE.variables


if docker ps | grep -q "rabbit-${PORT_BEGIN}"; then
  ../stop.sh
fi


if docker ps -a | grep -q "rabbit-${PORT_BEGIN}"; then
  for port in `seq $PORT_BEGIN $PORT_END`; do \
    docker container rm -f rabbit-${port}; \
  done
fi

if docker ps -a | grep -q "${RABBITMQ_HAPROXY}"; then
  docker container rm -f ${RABBITMQ_HAPROXY}
fi


if [ ! -d "./${PORT_BEGIN}" ]; then
  exit 0
fi

for port in `seq $PORT_BEGIN $PORT_END`; do \
  sudo rm -rf ./${port}; \
done

sudo rm -rf ./haproxy
