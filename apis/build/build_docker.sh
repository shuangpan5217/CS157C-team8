#!/bin/bash

cmd="docker run -p 9042:9042 -d --name cassandra cassandra"
printf "\n => $cmd\n\n"
eval $cmd

cmd="docker restart cassandra"
printf "\n => $cmd\n\n"
eval $cmd
sleep 5

cmd="docker build -t secretbox-api:v1 ."
printf "\n => $cmd\n\n"
eval $cmd
sleep 5

cmd="docker run -p 3000:3000 -d -e CASSANDRA_URL=cassandra:9042 --link=cassandra secretbox-api:v1"
printf "\n => $cmd\n\n"
eval $cmd
