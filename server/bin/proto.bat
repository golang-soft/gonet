protoc.exe  --go_out=../cmessage --proto_path=../cmessage  ../cmessage/*.proto
protoc.exe  --go_out=plugins=grpc:../smessage --proto_path=../smessage  ../smessage/*.proto