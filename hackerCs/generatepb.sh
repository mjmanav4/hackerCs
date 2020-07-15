 #!/usr/bin/env bash
 rm ./feed/feed.pb.go
 rm ./feed/feed.pb.gw.go

echo $PWD
protoc -I $PWD/feed -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./feed $PWD/feed/feed.proto
protoc -I/usr/local/include -I/$PWD/feed -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,allow_delete_body=true:./feed $PWD/feed/feed.proto
