#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables && source ../.env


node=${NODE_END}
if [ "$#" -gt 0 ]; then
  node=$1
fi


if ! docker ps | grep -q "rabbit-${node}"; then
  exit 0
fi


#docker exec -it rabbit-${node} bash
docker run -it --rm \
  --network=${BACKEND_NETWORK_NAME} \
  -v $(pwd)/scripts:/scripts \
  -v $(pwd)/build/${node}:/var/lib/rabbitmq \
  -e RABBITMQ_NODENAME=rabbit@rabbit-${node} \
  ${MANAGEMENT_IMAGE_FULL_NAME} \
  bash
