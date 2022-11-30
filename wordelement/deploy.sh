#!/bin/bash

#curl -fsSL get.docker.com -o get-docker.sh
# sh get-docker.sh

docker ps --filter="ancestor=bonggeek/elementservice" -q | xargs docker stop
docker pull bonggeek/elementservice
docker run -d --restart="always" -p 8080:8080 bonggeek/elementservice