#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables && source ../.env


if docker ps -a | grep -q "rabbit-${PORT_BEGIN}"; then
  exit 0
fi


./0-prepare.sh
../network-setup.sh


for node in `seq $PORT_END -1 $PORT_BEGIN`; do \
  docker run -d \
  --name rabbit-${node} \
  --network=${BACKEND_NETWORK_NAME} --hostname=rabbit-${node} \
  -v $(pwd)/build/${node}:/var/lib/rabbitmq \
  -e RABBITMQ_ERLANG_COOKIE=${ERLANG_COOKIE} \
  -e RABBITMQ_NODENAME=rabbit@rabbit-${node} \
  ${MANAGEMENT_IMAGE_FULL_NAME}; \
done
