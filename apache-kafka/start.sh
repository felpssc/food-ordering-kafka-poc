#!/bin/bash

docker-compose start zookeeper
sleep 3
docker-compose start kafka
sleep 5
docker-compose start kafka-topics-generator
sleep 2
docker-compose start control-center
sleep 2
docker-compose start kafka-connect
sleep 2
sudo sysctl -w vm.max_map_count=262144
sleep 1
docker-compose start es01
sleep 2
docker-compose start kibana
