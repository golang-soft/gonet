FROM golang:latest
#RUN apk update && apk add gcc git
WORKDIR /xxx
RUN ls
COPY . .
RUN pwd
RUN ls
WORKDIR /xxx/server
#RUN go get github.com/coreos/go-systemd/journal
RUN go build ./...
#RUN ["/bin/sh", "./bin/build_docker.sh"]

#FROM golang:latest
#ADD . /go/
#RUN ls
##COPY . /go/
#WORKDIR $GOPATH/server
#RUN ls
#WORKDIR /go/
#RUN pwd
##RUN ls server
#RUN ["/bin/sh", "./server/bin/build_docker.sh"]
##RUN ./build.bat
##RUN chmod +x server
##RUN pwd
#WORKDIR /go/server
#RUN rm $GOPATH/go.mod
#RUN /usr/local/go/bin/go build .
#
#EXPOSE 8081 31200 31300 31700
##ENTRYPOINT  ["./server", "netgate"]
#WORKDIR /go/server/bin
#RUN pwd
#RUN ls
#CMD ["./server.exe", "account"]
#CMD ["./server.exe", "world"]
#CMD ["./server.exe", "netgate"]

#USER root



#FROM centos:latest
#COPY ./bin /usr/local/bin
#ENV GATEWAY_LOG_LEVEL=info
#EXPOSE 8081 31100 31200 31300 31700
#WORKDIR /usr/local/bin
#RUN chmod u+x server
#RUN chmod u+x start.sh
#ENTRYPOINT  ["/bin/sh", "./start.sh"]
#USER root

#docker文件运行
#docker build -t dockerfile .
#删除空的镜像
#docker rmi -f $(docker images | grep "none" | awk '{print $3}')
#后台运行容器
#docker run -d -t -p31700:31700 -p31100:31100 -p31700:31700 dockerfile
#docker run -i -t dockerfile