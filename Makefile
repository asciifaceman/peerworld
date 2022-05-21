.PHONY: protos

#protoc -I service/ service/*.proto --go_out=plugins=grpc:service

protos: ## Compiles protos
	@echo "Building protos ..."
	protoc -I pkg/proto pkg/proto/*.proto --go_out=plugins=grpc:pkg/proto
