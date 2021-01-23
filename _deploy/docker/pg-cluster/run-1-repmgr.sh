#!/usr/bin/env bash

cd $(dirname $0) \
&& source ./NODE.variables


if docker ps | grep -q "pg-${NODE_BEGIN}"; then
  exit 0
fi

if ! docker ps -a | grep -q "pg-${NODE_BEGIN}"; then
  exit 1
fi


for node in `seq $NODE_BEGIN $NODE_END`; do \
  docker start pg-${node}; \
done
