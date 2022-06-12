#!/bin/bash
echo "now dir: "$(pwd)

# build image
cd ../
docker build . -t httpserver:v1.1
if [ 0 != $? ]
then
  echo "docker build faild"
  exit 1
fi

# deployment
cd deployment/template
docker images
