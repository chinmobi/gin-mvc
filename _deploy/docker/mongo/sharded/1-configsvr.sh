#!/usr/bin/env bash

cd $(dirname $0) && source ./NODE.variables \
&& source ../DB.variables \
&& source ../.env


set_name=${CONFIGSVR_REPLICA_SET_NAME}

BEGIN=${CONFIGSVR_NODE_BEGIN}
END=${CONFIGSVR_NODE_END}

primary_name="${set_name}-${BEGIN}"

if docker ps -a | grep -q "${primary_name}"; then
  exit 0
fi


./0-prepare.sh


# --- primary ---

docker run -d \
  --name ${primary_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${primary_name} \
  -v $(pwd)/build/${primary_name}:/bitnami \
  -e MONGODB_SHARDING_MODE=configsvr \
  -e MONGODB_REPLICA_SET_MODE=primary \
  -e MONGODB_REPLICA_SET_NAME=${set_name} \
  -e MONGODB_PORT_NUMBER=${PORT} \
  -e MONGODB_ADVERTISED_HOSTNAME=${primary_name} \
  -e MONGODB_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${SHARDED_IMAGE_FULL_NAME}


# --- secondaries ---

BEGIN=$(($BEGIN+1))

for node in `seq $BEGIN $END`; do
  node_name="${set_name}-${node}"

  docker run -d \
  --name ${node_name} \
  --network=${BACKEND_NETWORK_NAME} --hostname=${node_name} \
  -v $(pwd)/build/${node_name}:/bitnami \
  -e MONGODB_SHARDING_MODE=configsvr \
  -e MONGODB_REPLICA_SET_MODE=secondary \
  -e MONGODB_REPLICA_SET_NAME=${set_name} \
  -e MONGODB_PORT_NUMBER=${PORT} \
  -e MONGODB_ADVERTISED_HOSTNAME=${node_name} \
  -e MONGODB_PRIMARY_HOST=${primary_name} \
  -e MONGODB_PRIMARY_PORT_NUMBER=${PORT} \
  -e MONGODB_PRIMARY_ROOT_PASSWORD=${MONGODB_ROOT_PASSWORD} \
  -e MONGODB_REPLICA_SET_KEY=${REPLICA_SET_KEY} \
  ${SHARDED_IMAGE_FULL_NAME}
done
