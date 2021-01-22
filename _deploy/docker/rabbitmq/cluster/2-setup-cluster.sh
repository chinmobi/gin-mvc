#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables


if docker ps | grep -q "${RABBITMQ_HAPROXY}"; then
  exit 0
fi

if ! docker ps | grep -q "rabbit-${PORT_BEGIN}"; then
  exit 1
fi


STOP_CMD="rabbitmqctl stop_app"
RESET_CMD="rabbitmqctl reset"
JOIN_CMD="rabbitmqctl join_cluster rabbit@rabbit-${PORT_END}"
JOIN_RAM_CMD="rabbitmqctl join_cluster --ram rabbit@rabbit-${PORT_END}"
START_CMD="rabbitmqctl start_app"


END=$(($PORT_END-1))

CTL_CMD="${STOP_CMD} && ${RESET_CMD} && ${JOIN_CMD} && ${START_CMD}"

docker exec -it rabbit-${END} \
bash -c "${CTL_CMD}"


END=$(($END-1))

CTL_CMD="${STOP_CMD} && ${RESET_CMD} && ${JOIN_RAM_CMD} && ${START_CMD}"

for node in `seq $END -1 $PORT_BEGIN`; do
  docker exec -it rabbit-${node} \
  bash -c "${CTL_CMD}"
done
