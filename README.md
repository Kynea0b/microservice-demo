# microservice-demo
microservice sample project.use Go, kubernetes, gRPC...

## microservice

### grpc

#### module

```
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

#### generate grpc code

```
mkdir -p proto/grpc
make proto
```

