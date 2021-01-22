#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


port=${PORT_BEGIN}
if [ "$#" -gt 0 ]; then
  port=$1
fi


if ! docker ps | grep -q "redis-${port}"; then
  exit 0
fi


docker exec -it redis-${port} \
sh -c "redis-cli -c -p ${port}"
