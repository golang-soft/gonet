#!/bin/sh
./server "center" &
./server "worlddb" &

./server "world" &
#sleep 3
./server "world" &
#sleep 1
./server "account" &
#sleep 2
./server "account" &
#sleep 1
./server "netgate" &
#sleep 3
./server "netgate" &
#sleep 3
./server "netgate" &
#sleep 1
./server "login" &
./server "zone" &
./server "grpcserver" &