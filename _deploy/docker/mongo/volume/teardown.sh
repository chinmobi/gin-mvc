#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


if docker volume ls | grep -q "${VOLUME_NAME}"; then
  docker volume rm -f ${VOLUME_NAME}
fi
