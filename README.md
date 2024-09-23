# go-device-microservice-proto

## PROTO GENERATE

```
protoc --go_out=. --go_opt=Mprotos/device.proto=. --go-grpc_out=. --go-grpc_opt=Mprotos/device.proto=. protos/device.proto
```