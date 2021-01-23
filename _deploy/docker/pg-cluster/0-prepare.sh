#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


../networks/setup.sh


if [ -d "./build/pg-${NODE_BEGIN}" ]; then
  exit 0
fi


for node in `seq $NODE_BEGIN $NODE_END`; do \
  mkdir -p ./build/pg-${node} \
  && chmod a+w ./build/pg-${node}; \
done
