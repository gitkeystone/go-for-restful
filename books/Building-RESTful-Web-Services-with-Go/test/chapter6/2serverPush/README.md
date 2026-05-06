# gRPC
```bash
// 在 protofiles 目录执行
protoc "transaction.proto" \
  --go_out=.  --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative
```