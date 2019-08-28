#!/bin/sh

docker build -t computer-api . 
docker run -d -p 1337:1337 --rm computer-api

