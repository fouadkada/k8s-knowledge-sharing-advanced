current_dir = $(shell pwd)

build: ## generate go code from protobuf message
	docker run --rm -v $(current_dir):$(current_dir) -w $(current_dir) znly/protoc -I pkg/proto --go_out=plugins=grpc:pkg/proto pkg/proto/recommendation.proto
	docker run --rm -v $(current_dir):$(current_dir) -w $(current_dir) sagadash/protoc-gen-grpc-web -I pkg/proto --js_out=import_style=commonjs,binary:pkg/proto --grpc_out=grpc_js:pkg/proto pkg/proto/recommendation.proto