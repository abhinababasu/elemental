#!/bin/bash

VERSION=0.1
CONTAINER=elementservice

echo Build go sources
go build ./wordelementservice.go

echo Building Image $CONTAINER:$VERSION
docker build -t $CONTAINER:$VERSION .
docker tag $CONTAINER:$VERSION $DOCKER_ID_USER/$CONTAINER

echo Done ................................
echo 
echo Run Container using
echo docker run -d -p 8080:8080 $CONTAINER:$VERSION

