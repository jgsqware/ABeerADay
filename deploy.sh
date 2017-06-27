#!/bin/bash +eu

docker login -u jgsqware -p ${DOCKER_PASSWORD} -e " " quay.io
docker tag quay.io/jgsqware/${PROJECT_NAME}:$(git log --format="%H" -n 1) quay.io/jgsqware/${PROJECT_NAME}:latest
docker push quay.io/jgsqware/${PROJECT_NAME}:$(git log --format="%H" -n 1)
docker push quay.io/jgsqware/${PROJECT_NAME}:latest
sudo ./clairctl analyze quay.io/jgsqware/${PROJECT_NAME}:$(git log --format="%H" -n 1)
sudo chown -R ubuntu:ubuntu /home/ubuntu/.kube
sudo chown -R $USER /home/ubuntu/.config
sed -i 's/{{SHA1}}/'$(git log --format="%H" -n 1)'/' kubernetes/deployment.yaml
kubectl apply -f kubernetes/deployment.yaml
