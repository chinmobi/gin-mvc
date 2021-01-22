#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if [ -d "./build/${PORT_BEGIN}" ]; then
  exit 0
fi


for port in `seq $PORT_BEGIN $PORT_END`; do \
  mkdir -p ./build/${port}; \
done

mkdir -p ./build/haproxy
