#!/bin/bash

# curl -fsSL get.docker.com -o get-docker.sh
# sh get-docker.sh

docker ps -q | xargs docker stop 
docker pull bonggeek/elementservice
docker run -d --restart="always" -p 80:8080 bonggeek/elementservice
