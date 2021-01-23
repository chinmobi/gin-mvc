#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables && source ../.env


if docker ps | grep -q "${RABBITMQ_HAPROXY}"; then
  exit 0
fi

if ! docker ps | grep -q "rabbit-${NODE_BEGIN}"; then
  exit 1
fi


### Join as cluster nodes

UNSET_COOKIE="unset RABBITMQ_ERLANG_COOKIE"
STOP_CMD="rabbitmqctl stop_app"
RESET_CMD="rabbitmqctl reset"
JOIN_CMD="rabbitmqctl join_cluster rabbit@rabbit-${NODE_END}"
JOIN_RAM_CMD="rabbitmqctl join_cluster --ram rabbit@rabbit-${NODE_END}"
START_CMD="rabbitmqctl start_app"


END=$(($NODE_END-1))

CTL_CMD="${UNSET_COOKIE} && ${STOP_CMD} && ${RESET_CMD} && ${JOIN_CMD} && ${START_CMD}"

docker exec -it rabbit-${END} \
bash -c "${CTL_CMD}"


END=$(($END-1))

CTL_CMD="${UNSET_COOKIE} && ${STOP_CMD} && ${RESET_CMD} && ${JOIN_RAM_CMD} && ${START_CMD}"

for node in `seq $END -1 $NODE_BEGIN`; do
  docker exec -it rabbit-${node} \
  bash -c "${CTL_CMD}"
done


### Set cluster policy (Mirroring)

docker run -it --rm \
  --network=${BACKEND_NETWORK_NAME} \
  -v $(pwd)/scripts:/scripts \
  -v $(pwd)/build/${NODE_END}:/var/lib/rabbitmq \
  -e RABBITMQ_NODENAME=rabbit@rabbit-${NODE_END} \
  ${MANAGEMENT_IMAGE_FULL_NAME} \
  /scripts/set-policy.sh
