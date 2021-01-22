#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


port=${PORT_END}
if [ "$#" -gt 0 ]; then
  port=$1
fi


if ! docker ps | grep -q "rabbit-${port}"; then
  exit 0
fi


docker exec -it rabbit-${port} bash


### Check cluster status
# rabbitmqctl cluster_status

### Set as mirror cluster
# rabbitmqctl set_policy ha-all "^" '{"ha-mode":"all"}'
