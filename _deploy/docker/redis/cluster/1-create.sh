#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables && source ../.env


if docker ps -a | grep -q "redis-${PORT_BEGIN}"; then
  exit 0
fi


./0-prepare.sh
../network-setup.sh


for port in `seq $PORT_BEGIN $PORT_END`; do \
  docker run -d \
  --name redis-${port} \
  --network=${BACKEND_NETWORK_NAME} --hostname=redis-${port} \
  -p ${port}:${port} -p 1${port}:1${port} \
  -v $(pwd)/build/${port}/conf/redis.conf:/usr/local/etc/redis/redis.conf \
  -v $(pwd)/build/${port}/data:/data \
  --sysctl net.core.somaxconn=1024 \
  ${IMAGE_FULL_NAME} \
  redis-server /usr/local/etc/redis/redis.conf; \
done


function get_container_ip() {
  docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-$1
}


REDIS_CLI_CMD="redis-cli --cluster create"

for port in `seq $PORT_BEGIN $PORT_END`; do \
  REDIS_CLI_CMD="${REDIS_CLI_CMD} $(get_container_ip ${port}):${port}"; \
done

REDIS_CLI_CMD="${REDIS_CLI_CMD} --cluster-replicas 1"


docker exec -it redis-${PORT_BEGIN} \
sh -c "echo '${REDIS_CLI_CMD}' && ${REDIS_CLI_CMD}"
