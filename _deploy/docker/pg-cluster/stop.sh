#!/usr/bin/env bash

cd $(dirname $0) \
&& source ./NODE.variables \
&& source ./.env


if ! docker ps | grep -q "pg-${NODE_END}"; then
  exit 0
fi


docker stop ${PGPOOL_CONTAINER_NAME}

for node in `seq $NODE_END -1 $NODE_BEGIN`; do \
  docker stop pg-${node}; \
done
