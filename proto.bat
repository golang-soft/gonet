protoc.exe --plugin=protoc-gen-go=protoc-gen-go.exe  --go_out=../message  --proto_path=../message	
::protoc.exe --js_out=../message  --proto_path=../message	
::protoc.exe --cpp_out=../message/c++  --proto_path=../message	
::protoc.exe -o ../message/pb/client.pb --proto_path=../message	
