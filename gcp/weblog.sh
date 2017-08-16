#!/bin/bash

NAME=`kubectl get pod -l app=web | tail -n 1 | awk '{print $1}'`
#echo $NAME

echo 'kubectl describe pod ===================================>'
kubectl describe pod ${NAME}

echo ''
echo ''
echo 'kubectl logs ===================================>'
kubectl logs ${NAME}
