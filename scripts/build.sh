#!/bin/bash

# this scripts builds and tags with docker

# it takes one argument, the tag version

TAG=latest

if [ $# -eq 1 ]; then
    TAG=$1
fi

docker build -t dekuyo/chatmap-server:$TAG -f ./Dockerfile --target production .
docker push dekuyo/chatmap-server:$TAG
