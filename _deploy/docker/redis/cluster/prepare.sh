#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if [ -d "./_deploy/${PORT_BEGIN}/conf" ]; then
  exit 0
fi


for port in `seq $PORT_BEGIN $PORT_END`; do \
  mkdir -p ./_deploy/${port}/conf \
  && PORT=${port} envsubst < ./node-conf.tmpl > ./_deploy/${port}/conf/redis.conf \
  && mkdir -p ./_deploy/${port}/data; \
done
