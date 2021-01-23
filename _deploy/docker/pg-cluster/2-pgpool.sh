#!/usr/bin/env bash

cd $(dirname $0) \
&& source ./PGSQL.variables \
&& source ./PGPOOL.variables \
&& source ./NODE.variables \
&& source ./.env


if docker ps | grep -q "${PGPOOL_CONTAINER_NAME}"; then
  exit 0
fi

if ! docker ps | grep -q "pg-${NODE_BEGIN}"; then
  exit 1
fi


backend_nodes="0:pg-${NODE_BEGIN}:${PORT}"

BEGIN=$(($NODE_BEGIN+1))
node_index=1
for node in `seq $BEGIN $NODE_END`; do
  backend_nodes="${backend_nodes},${node_index}:pg-${node}:${PORT}"
  node_index=$(($node_index+1))
done


docker run -d \
  --name ${PGPOOL_CONTAINER_NAME} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${PGPOOL_CONTAINER_NAME} \
  -p ${HOST_PORT}:${PORT} \
  --env PGPOOL_BACKEND_NODES=${backend_nodes} \
  --env PGPOOL_ENABLE_LOAD_BALANCING=${PGPOOL_ENABLE_LOAD_BALANCING} \
  --env PGPOOL_ENABLE_LDAP=${PGPOOL_ENABLE_LDAP} \
  --env PGPOOL_SR_CHECK_USER=${POSTGRESQL_USERNAME} \
  --env PGPOOL_SR_CHECK_PASSWORD=${POSTGRESQL_PASSWORD} \
  --env PGPOOL_POSTGRES_USERNAME=${PGPOOL_POSTGRES_USERNAME} \
  --env PGPOOL_POSTGRES_PASSWORD=${POSTGRESQL_POSTGRES_PASSWORD} \
  --env PGPOOL_ADMIN_USERNAME=${PGPOOL_ADMIN_USERNAME} \
  --env PGPOOL_ADMIN_PASSWORD=${PGPOOL_ADMIN_PASSWORD} \
  ${PGPOOL_IMAGE_FULL_NAME}
