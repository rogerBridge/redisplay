#!/usr/bin/env bash
go build -o redisBuy *.go
docker stop redisshop && docker rm redisshop
docker rmi leo2n/redisshop:test
docker build -t leo2n/redisshop:test .
docker run -d --name redisshop -p 4000:4000 --network=redisStore leo2n/redisshop:test
