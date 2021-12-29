protoc.exe  --go_out=../message --proto_path=../message  ../message/*.proto
protoc.exe  --go_out=plugins=grpc:../smessage --proto_path=../smessage  ../smessage/*.proto