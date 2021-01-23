#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


if ! docker images | grep -q "${PGSQL_IMAGE_REPOSITORY}"; then
  docker pull ${PGSQL_IMAGE_FULL_NAME}
fi

if ! docker images | grep -q "${REPMGR_IMAGE_REPOSITORY}"; then
  docker pull ${REPMGR_IMAGE_FULL_NAME}
fi

if ! docker images | grep -q "${PGPOOL_IMAGE_REPOSITORY}"; then
  docker pull ${PGPOOL_IMAGE_FULL_NAME}
fi
