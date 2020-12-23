## GRPC

ืทำสอบ gRPC ไม่ได้ celan code นะจ้ะ

### Install compiler Protobuf

```bash
$ brew install protobuf
$ protoc --version

libprotoc 3.14.0
$ brew install protoc-gen-go
```

### Generate code by protoc & protoc-gen-go in protos directory

```bash
protoc -I . \
   --go_out . --go_opt paths=source_relative \
   --go-grpc_out . --go-grpc_opt paths=source_relative \
   manga.proto
```

and Rest

```bash
protoc -I . --grpc-gateway_out . \
     --grpc-gateway_opt logtostderr=true \
     --grpc-gateway_opt paths=source_relative \
     manga.proto
```

### Let's do it
