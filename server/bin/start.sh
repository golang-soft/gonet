#!/bin/sh

./server "center" &
sleep 1
./server "worlddb" &
sleep 1
./server "world" &
sleep 1
./server "account" &
sleep 1
./server "netgate" &
sleep 1
./server "login" &
sleep 3
./server "zone" &
sleep 3
./server "grpcserver" &
sleep 8
./server "world" &
sleep 3
./server "account" &
sleep 3
./server "netgate" &
sleep 3
