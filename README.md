# microservice-demo
microservice sample project.use Go, kubernetes, gRPC...

## grpc

```
go get -u google.golang.org/grpc
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### generate grpc code

```
mkdir -p proto/grpc
make proto
```

## request

### draw gacha

```
curl -XPOST -H 'Content-Type: application/json' -d '{"user_id": 1}' localhost:50942/draw
```

### get gacha histories

```
curl localhost:50942/histories/1
```

### get item inventory

```
curl localhost:50942/inventories/1
```
