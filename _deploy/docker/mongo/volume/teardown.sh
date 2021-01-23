#!/usr/bin/env bash

cd $(dirname $0) && source ../.env


if [ -d "${HOME_MOUNT}" ]; then
  sudo rm -rf ${HOME_MOUNT}
fi

if docker volume ls | grep -q "${VOLUME_NAME}"; then
  docker volume rm -f ${CONFIG_VOLUME_NAME}
  docker volume rm -f ${VOLUME_NAME}
fi
