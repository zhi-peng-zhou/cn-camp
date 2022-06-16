#!/bin/bash
echo "now dir: "$(pwd)

# build image
cd ../
docker build . -t httpserver:v1.1
docker tag httpserver:v1.1 192.168.20.49:30002/myserver/httpserver:v1.1
docker push 192.168.20.49:30002/myserver/httpserver:v1.1
if [ 0 != $? ]
then
  echo "docker build faild"
  exit 1
fi

# deployment
kubectl get pod
kubectl apply -f template/config.yaml
kubectl apply -f template/serv-deployment.yaml  
