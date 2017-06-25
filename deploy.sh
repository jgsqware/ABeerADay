#!/bin/bash +eu

docker login -u juliengarcia -p ${DOCKER_PASSWORD} quay.io
docker push quay.io/${PROJECT_ID}/${PROJECT_NAME}:$CIRCLE_SHA1
docker push quay.io/${PROJECT_ID}/${PROJECT_NAME}:latest
sudo chown -R ubuntu:ubuntu /home/ubuntu/.kube
kubectl patch deployment ${PROJECT_NAME} -p '{"spec":{"template":{"spec":{"containers":[{"name":"${PROJECT_NAME}","image":"quay.io/'"${PROJECT_ID}"'/'"${PROJECT_NAME}"':'"$CIRCLE_SHA1"'"}]}}}}'
