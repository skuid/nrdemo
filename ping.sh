#!/bin/sh

apk -U add curl

while :
do
    sleep 1;
    curl http://demoapp:8080/ -s
done
