# protobuf

```bash
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 33.0+

# 安装插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 开始编译
protoc --go_out=. person.proto
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

# 在basicExample中添加模块
go mod edit --replace protofiles=../protofiles    
```