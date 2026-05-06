module github.com/gitkeystone/basic-example

go 1.26.1

replace protofiles => ../protofiles

require (
	github.com/golang/protobuf v1.5.0
	protofiles v0.0.0-00010101000000-000000000000
)

require google.golang.org/protobuf v1.36.11 // indirect
