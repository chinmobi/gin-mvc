#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables && source ../.env


port=${PORT_END}
if [ "$#" -gt 0 ]; then
  port=$1
fi


if ! docker ps | grep -q "rabbit-${port}"; then
  exit 0
fi


#docker exec -it rabbit-${port} bash
docker run -it --rm \
  --network=${BACKEND_NETWORK_NAME} \
  -v $(pwd)/scripts:/scripts \
  -v $(pwd)/build/${port}:/var/lib/rabbitmq \
  -e RABBITMQ_NODENAME=rabbit@rabbit-${port} \
  ${MANAGEMENT_IMAGE_FULL_NAME} \
  bash
