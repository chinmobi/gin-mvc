#!/usr/bin/env bash

cd $(dirname $0) \
&& source ./PGSQL.variables \
&& source ./NODE.variables \
&& source ./.env


if ! docker ps | grep -q "${PGPOOL_CONTAINER_NAME}"; then
  exit 0
fi


docker run -it --rm \
  --network=${BACKEND_NETWORK_NAME} \
  ${PGSQL_IMAGE_FULL_NAME} \
  psql -h ${PGPOOL_CONTAINER_NAME} \
  -U ${POSTGRESQL_USERNAME} \
  -d ${POSTGRESQL_DATABASE}
