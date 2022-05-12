#!/bin/bash

PROD_NAME=sphinx-proxy

make verify-build

kubectl delete deployment $PROD_NAME -n kube-system

minikube ssh "cd /home/coast/dev_box/$PROD_NAME/output/linux/amd64/&&docker build -f Dockerfile -t test/$PROD_NAME ."

kubectl apply -f ./output/linux/amd64/01-$PROD_NAME.yaml