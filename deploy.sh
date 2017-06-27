#!/bin/bash 
set -euo pipefail

# sudo chown -R ubuntu:ubuntu /home/ubuntu/.kube
sudo chown -R $USER /home/ubuntu/.config
sed -i 's/{{SHA1}}/'$(git log --format="%H" -n 1)'/' kubernetes/deployment.yaml
kubectl apply -f kubernetes/deployment.yaml
