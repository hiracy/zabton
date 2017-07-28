#!/bin/sh

set -e
PATH=/usr/local/bin:/usr/sbin:/sbin:/usr/bin:/bin

ZABBIX_DB_DOCKER_IMAGE="monitoringartist/zabbix-db-mariadb"
ZABBIX_DOCKER_IMAGE="monitoringartist/zabbix-xxl:latest"
ZABBIX_DB_DOCKER_CONTAINER_NAME=(`echo "${ZABBIX_DOCKER_IMAGE}/zabbix-db" | sed -e 's/\//-/g' | sed -e 's/\:/-/g'`)
ZABBIX_CONTAINER_NAME=(`echo "${ZABBIX_DOCKER_IMAGE}" | sed -e 's/^.*\///g' | sed -e 's/\:/-/g'`)
DOCKER_CACHE_DIR="/var/lib/docker/cache"
DOCKER_CACHE_ZABBIX_DIR="${DOCKER_CACHE_DIR}/${PROJECT_NAME}"
DOCKER_CACHE_ZABBIX_DB_DIR="${DOCKER_CACHE_DIR}/${PROJECT_NAME}/db"
DOCKER_CACHE_IMAGE_PATH_ZABBIX_DB="${DOCKER_CACHE_ZABBIX_DB_DIR}/docker.tar"
ZABBIX_CACHE_IMAGE_DIR="${DOCKER_CACHE_ZABBIX_DIR}/${ZABBIX_CONTAINER_NAME}"
ZABBIX_CACHE_IMAGE_PATH="${DOCKER_CACHE_ZABBIX_DIR}/${ZABBIX_CONTAINER_NAME}/zabbix_server.tar"
CLEAN=0
PROJECT_NAME=$(cd $(dirname $0)/.. && basename `pwd`)
ZABBIX_DB_USER="zabbix"
ZABBIX_DB_PASSWORD="zabbix"
ZABBIX_HTTP_PROXY_PORT=8080

usage () {
  echo "Usage: `basename $0` [options]"
  echo "[-h|--help ]             :show help. "
  echo "[-i|--image]             :target docker image"
  echo "[-d|--docker-cache-dir]  :docker image cache path"
  echo "[-p|--http-port]         :zabbix http proxy port"
  echo "[-c|--clean]             :clean docker processes"

  exit 0
}

parse_opts () {
  while [ $# -ne 0 ]
  do
    case "$1" in
      -i|--image)             ZABBIX_DOCKER_IMAGE=$2; shift 2 ;;
      -d|--docker-cache-dir)  DOCKER_CACHE_DIR=$2; shift 2 ;;
      -p|--http-port)         ZABBIX_HTTP_PROXY_PORT=$2; shift 2 ;;
      -c|--clean)             CLEAN=1; shift ;;
      --) shift; break ;;
      *) usage;;
    esac
  done
}

parse_opts $@

if ! sudo docker info > /dev/null 2>&1; then
  echo "docker daemon is not running."
  exit 0
fi

if [ ${CLEAN} -eq 1 ];then
  set +e
  if sudo docker ps -a | grep -w "${ZABBIX_DB_DOCKER_IMAGE}" > /dev/null 2>&1; then
    sudo docker ps -a | grep -w "${ZABBIX_DB_DOCKER_IMAGE}" | awk '{print $1}' | xargs sudo docker rm -f
  fi

  if sudo docker ps -a | grep -w "${ZABBIX_DOCKER_IMAGE}" > /dev/null 2>&1; then
    sudo docker ps -a | grep -w "${ZABBIX_DOCKER_IMAGE}" | awk '{print $1}' | xargs sudo docker rm -f
  fi

  if [ -d "${ZABBIX_CACHE_IMAGE_DIR}" ];
  then
    rm -fr ${ZABBIX_CACHE_IMAGE_DIR}
  fi

  exit 0
fi

if file ${DOCKER_CACHE_IMAGE_PATH_ZABBIX_DB} | grep empty; then
  sudo docker load --input ${DOCKER_CACHE_IMAGE_PATH_ZABBIX_DB}
fi

if ! sudo docker ps -a | grep -w "${ZABBIX_DB_DOCKER_IMAGE}" > /dev/null 2>&1; then
  echo "run docker ${ZABBIX_DB_DOCKER_IMAGE}"

  sudo docker run \
    -d \
    --name ${ZABBIX_DB_DOCKER_CONTAINER_NAME} \
    --env="MARIADB_USER=${ZABBIX_DB_USER}" \
    --env="MARIADB_PASS=${ZABBIX_DB_PASSWORD}" \
    ${ZABBIX_DB_DOCKER_IMAGE}

  set +e
  while :
  do
    sudo docker logs ${ZABBIX_DB_DOCKER_CONTAINER_NAME} | grep "You can now connect to this MariaDB Server using" > /dev/null 2>&1
    if [ $? -eq 0 ]; then
      break
    fi
    echo -en "."

    sleep 1
  done

  echo -en "\n"
  set -e
fi

if [ ! -f "${DOCKER_CACHE_IMAGE_PATH_ZABBIX_DB}" ]; then
  mkdir -p ${DOCKER_CACHE_ZABBIX_DB_DIR}
  echo "saving ${ZABBIX_DB_DOCKER_IMAGE} > ${DOCKER_CACHE_IMAGE_PATH_ZABBIX_DB} ..."
  sudo docker save "${ZABBIX_DB_DOCKER_IMAGE}" > ${DOCKER_CACHE_IMAGE_PATH_ZABBIX_DB}
fi

if ! sudo docker ps -a | grep -w "${ZABBIX_DOCKER_IMAGE}" > /dev/null 2>&1; then
  echo "run zabbix docker ${ZABBIX_DOCKER_IMAGE}"

  if file "${ZABBIX_CACHE_IMAGE_PATH}" | grep empty; then
    sudo docker load --input ${ZABBIX_CACHE_IMAGE_PATH}
  fi

  sudo docker run \
    -d \
    --name "${ZABBIX_CONTAINER_NAME}" \
    -p ${ZABBIX_HTTP_PROXY_PORT}:80 \
    -v /etc/localtime:/etc/localtime:ro \
    --link ${ZABBIX_DB_DOCKER_CONTAINER_NAME}:zabbix.db \
    --env="ZS_DBHost=zabbix.db" \
    --env="ZS_DBUser=${ZABBIX_DB_USER}" \
    --env="ZS_DBPassword=${ZABBIX_DB_PASSWORD}" \
    ${ZABBIX_DOCKER_IMAGE}

  set +e
  while :
  do
#    sudo docker exec -it ${ZABBIX_CONTAINER_NAME}  ps -C zabbix_server | grep "API is available" > /dev/null 2>&1
    sudo docker exec -it ${ZABBIX_CONTAINER_NAME}  ps -C zabbix_server > /dev/null 2>&1

    if [ $? -eq 0 ]; then
      break
    fi
    echo -en "."

    sleep 1
  done

  echo -en "\n"
  set -e
fi

if [ ! -f "${ZABBIX_CACHE_IMAGE_PATH}" ]; then
  mkdir -p ${ZABBIX_CACHE_IMAGE_DIR}
  echo "saving ${ZABBIX_DOCKER_IMAGE} > ${ZABBIX_CACHE_IMAGE_PATH} ..."
  sudo docker save "${ZABBIX_DOCKER_IMAGE}" > ${ZABBIX_CACHE_IMAGE_PATH}
fi
