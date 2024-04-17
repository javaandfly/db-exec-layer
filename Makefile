bindir:
	mkdir -p bin/

linux: bindir linux-db-server 

linux-db-server:
	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -ldflags='-extldflags="-static"' -tags 'osusergo netgo static_build' -o bin/wallet cmd/main.go 


protoc:  dbservice

dbservice:
	protoc  --proto_path=$(GOPATH)/src/  --proto_path=. --go_out=./protocol/pb/  --go-grpc_out=./protocol/pb/ --validate_out="lang=go:./protocol/pb/"  protocol/pb/dbservice_protofile/dbservice.proto 




