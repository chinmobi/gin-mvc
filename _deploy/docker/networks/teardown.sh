#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


NETWORK_NAMES=(${BACKEND_NETWORK_NAME} ${FRONTEND_NETWORK_NAME})

for NETWORK_NAME in ${NETWORK_NAMES[@]}; do

  if docker network ls | grep -q "${NETWORK_NAME}"; then
    docker network rm ${NETWORK_NAME}
  fi

done
