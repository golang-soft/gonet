#!/bin/sh
#sh stop.sh
#cd ./server/
go build
cp server/server ./bin
rm -rf server
#go install
#cd ./client
go build
cp server/client ./bin
rm -rf client
#go install
