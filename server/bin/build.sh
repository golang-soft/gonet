#!/bin/sh
sh stop.sh
cd ./../
go build
cp server ./bin
rm -rf server
#go install
cd ./client
go build
cp client ./../bin
rm -rf client
#go install