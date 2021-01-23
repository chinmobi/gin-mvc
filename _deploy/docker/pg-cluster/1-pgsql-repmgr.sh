#!/usr/bin/env bash

cd $(dirname $0) \
&& source ./PGSQL.variables \
&& source ./NODE.variables \
&& source ./.env


if docker ps -a | grep -q "pg-${NODE_BEGIN}"; then
  exit 0
fi


./0-prepare.sh


partner_nodes="pg-${NODE_BEGIN}"

BEGIN=$(($NODE_BEGIN+1))
for node in `seq $BEGIN $NODE_END`; do
  partner_nodes="${partner_nodes},pg-${node}"
done
partner_nodes="${partner_nodes}:${PORT}"


for node in `seq $NODE_BEGIN $NODE_END`; do \
  docker run -d \
  --name pg-${node} \
  --network=${BACKEND_NETWORK_NAME} --hostname=pg-${node} \
  -v $(pwd)/build/pg-${node}:/bitnami/postgresql \
  --env REPMGR_PARTNER_NODES=${partner_nodes} \
  --env REPMGR_NODE_NAME=pg-${node} \
  --env REPMGR_NODE_NETWORK_NAME=pg-${node} \
  --env REPMGR_PORT_NUMBER=${PORT} \
  --env REPMGR_PRIMARY_HOST=pg-${NODE_BEGIN} \
  --env REPMGR_PRIMARY_PORT=${PORT} \
  --env REPMGR_PASSWORD=${REPMGR_PASSWORD} \
  --env POSTGRESQL_POSTGRES_PASSWORD=${POSTGRESQL_POSTGRES_PASSWORD} \
  --env POSTGRESQL_USERNAME=${POSTGRESQL_USERNAME} \
  --env POSTGRESQL_PASSWORD=${POSTGRESQL_PASSWORD} \
  --env POSTGRESQL_DATABASE=${POSTGRESQL_DATABASE} \
  --env POSTGRESQL_NUM_SYNCHRONOUS_REPLICAS=1 \
  ${REPMGR_IMAGE_FULL_NAME}; \
done
