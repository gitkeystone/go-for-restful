```bash
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
protoc --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "encryption.proto"
```