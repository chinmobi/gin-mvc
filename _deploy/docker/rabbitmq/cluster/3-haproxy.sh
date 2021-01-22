#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables && source ../.env


if docker ps | grep -q "${RABBITMQ_HAPROXY}"; then
  exit 0
fi

if ! docker ps | grep -q "rabbit-${PORT_BEGIN}"; then
  exit 1
fi


HAPROXY_TAG=$(docker images | grep "${HAPROXY_REPOSITORY}" | awk -F ' ' '{print $2}')

if [ "x${HAPROXY_TAG}" = "x" ]; then
  exit 2
fi

HAPROXY_IMAGE="${HAPROXY_REPOSITORY}:${HAPROXY_TAG}"


amqp_server_nodes=""
mgmt_server_nodes=""

server_options="check inter 5s rise 2 fall 3"

for node in `seq $PORT_BEGIN $PORT_END`; do
  node_name="rabbit-${node}"

  amqp_server_nodes="${amqp_server_nodes}\\tserver  ${node_name} ${node_name}:${PORT}  ${server_options}\\n"
  mgmt_server_nodes="${mgmt_server_nodes}\\tserver  ${node_name} ${node_name}:${MGMT_PORT}\\n"
done

sed \
  -e "s/<mgmt_port>/$MGMT_PORT/" \
  -e "s/<mgmt_server_nodes>/$mgmt_server_nodes/" \
  -e "s/<amqp_port>/$PORT/" \
  -e "s/<amqp_server_nodes>/$amqp_server_nodes/" \
  -e "s/<admin_port>/$HAPROXY_ADMIN_PORT/" \
  "./haproxy-cfg.tmpl" \
  >"./build/haproxy/haproxy.cfg"


docker run -d \
  --name ${RABBITMQ_HAPROXY} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${RABBITMQ_HAPROXY} \
  -p ${HOST_PORT}:${PORT} -p ${MGMT_HOST_PORT}:${MGMT_PORT} \
  -p ${HAPROXY_ADMIN_HOST_PORT}:${HAPROXY_ADMIN_PORT} \
  -v $(pwd)/build/haproxy:/usr/local/etc/haproxy:ro \
  --user haproxy \
  ${HAPROXY_IMAGE}
