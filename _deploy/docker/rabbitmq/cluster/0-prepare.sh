#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if [ -d "./build/${NODE_BEGIN}" ]; then
  exit 0
fi


for node in `seq $NODE_BEGIN $NODE_END`; do \
  mkdir -p ./build/${node}; \
done

mkdir -p ./build/haproxy
