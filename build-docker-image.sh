#!/usr/bin/env bash

if [ ! -f "./Dockerfile" ]; then
  echo "··> [ERROR] Dockerfile file does not exist."
  exit 1
fi

echo "··> [INFO] Dependency [compile application] <··"
bash ./build.sh linux

if [ ! -f "./Dockerfile" ]; then
  echo "··> Dependency [compile application failed.] <··" && exit 1
fi
echo "··> [INFO] Dependency [compile application complete.] <··"
echo "··> [INFO] Start building docker image..."

docker build -t iluckin/bear:latest .

echo "··> [INFO] OK"

docker images | grep "iluckin/bear"
