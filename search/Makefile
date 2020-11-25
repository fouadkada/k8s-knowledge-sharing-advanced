current_dir = $(shell pwd)

build: ## generate go code from protobuf message
	docker run --rm -v $(current_dir):$(current_dir) -w $(current_dir) znly/protoc -I pkg/proto --go_out=plugins=grpc:pkg/proto pkg/proto/search.proto