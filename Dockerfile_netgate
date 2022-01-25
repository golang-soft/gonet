FROM golang:latest
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /xxx
RUN ls
COPY . .
RUN pwd
RUN ls
WORKDIR /xxx/server
RUN go get github.com/coreos/go-systemd/journal
RUN ls
RUN go build -o server .
RUN chmod +x server
RUN cp /xxx/server/server ./bin/
WORKDIR /xxx/server/bin
RUN pwd
RUN ls -al
EXPOSE 3000 31200 31300 31700 31599 3001
ENTRYPOINT  ["./server", "netgate"]
#ENTRYPOINT  ["/bin/sh", "./start.sh"]
#CMD ["./server", "account"]
#CMD ["./server", "world"]
#CMD ["./server", "netgate"]

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
#docker run -d -t -p3000:3000 dockerfile
#docker run -i -t dockerfile