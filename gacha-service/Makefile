.PHONY: proto
proto:
	protoc --proto_path=../ \
		--go_out=./pkg/grpc --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/grpc --go-grpc_opt=paths=source_relative \
		../proto/*.proto

.PHONY: draw
draw:
	grpcurl --plaintext localhost:8080 gacha.GachaService.Draw