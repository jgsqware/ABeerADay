#!/bin/bash +eu

docker login -u juliengarcia -p ${DOCKER_PASSWORD} -e garciagonzalez.julien@gmail.com quay.io
docker push quay.io/juliengarcia/${PROJECT_NAME}:$CIRCLE_SHA1
docker push quay.io/juliengarcia/${PROJECT_NAME}:latest
sudo chown -R ubuntu:ubuntu /home/ubuntu/.kube
sudo chown -R $USER /home/ubuntu/.config
kubectl patch deployment ${PROJECT_NAME} -p '{"spec":{"template":{"spec":{"containers":[{"name":"${PROJECT_NAME}","image":"quay.io/juliengarcia/'"${PROJECT_NAME}"':'"$CIRCLE_SHA1"'"}]}}}}'
