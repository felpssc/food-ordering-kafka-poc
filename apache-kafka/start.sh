#!/bin/bash

docker-compose restart zookeeper
sleep 3
docker-compose restart kafka
sleep 5
docker-compose restart kafka-topics-generator
sleep 2
docker-compose restart control-center
sleep 2
docker-compose restart kafka-connect
sleep 2
sudo sysctl -w vm.max_map_count=262144
sleep 1
docker-compose restart es01
sleep 2
docker-compose restart kibana
