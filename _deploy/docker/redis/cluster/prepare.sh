#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if [ -d "./build/${PORT_BEGIN}/conf" ]; then
  exit 0
fi


for port in `seq $PORT_BEGIN $PORT_END`; do \
  mkdir -p ./build/${port}/conf \
  && PORT=${port} envsubst < ./node-conf.tmpl > ./build/${port}/conf/redis.conf \
  && mkdir -p ./build/${port}/data; \
done
