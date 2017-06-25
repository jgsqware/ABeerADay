#!/bin/bash +eu

sudo /opt/google-cloud-sdk/bin/gcloud docker -- push us.gcr.io/${PROJECT_ID}/${PROJECT_NAME}
sudo chown -R ubuntu:ubuntu /home/ubuntu/.kube
kubectl patch deployment ${PROJECT_NAME} -p '{"spec":{"template":{"spec":{"containers":[{"name":"${PROJECT_NAME}","image":"us.gcr.io/'"${PROJECT_ID}"'/'"${PROJECT_NAME}"':'"$CIRCLE_SHA1"'"}]}}}}'
