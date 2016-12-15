#!/bin/bash

go run gracefulstoppablehttpd.go &
sleep 10

echo "" > ./gracelog.txt

for i in {1..10}; do wget -q -O- http://localhost:8180/0/$i >> ./gracelog.txt & done;

sleep 1

for i in {1..10}; do wget -q -O- http://localhost:8180/1/$i >> ./gracelog.txt & done;

go run gracefulstoppablehttpd.go &

sleep 1

curl http://localhost:8180/stop --insecure 

for i in {1..10}; do wget -q -O- http://localhost:8180/2/$i >> ./gracelog.txt & done;

curl http://localhost:8180/stop --insecure

sleep 10

echo $(grep "Hello World" ./gracelog.txt | wc -l) | tr -d '[[:space:]]'

echo " hits "