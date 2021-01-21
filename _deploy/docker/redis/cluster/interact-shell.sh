#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if ! docker ps | grep -q "redis-${PORT_BEGIN}"; then
  exit 0
fi


docker exec -it redis-${PORT_BEGIN} \
sh -c "redis-cli -c -p ${PORT_BEGIN}"
