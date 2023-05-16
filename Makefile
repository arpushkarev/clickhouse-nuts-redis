PHONY: generate
generate:
		mkdir -p pkg/item_v1
		protoc  --proto_path vendor.protogen --proto_path api/item_v1 \
				--go_out=pkg/item_v1 --go_opt=paths=source_relative \
				--go-grpc_out=pkg/item_v1 --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=pkg/item_v1 \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=source_relative \
				--validate_out lang=go:pkg/item_v1 --validate_opt=paths=source_relative \
				--swagger_out=allow_merge=true,merge_file_name=api:pkg/item_v1 \
				api/item_v1/service.proto

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google/protobuf ]; then \
			git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
			mkdir -p  vendor.protogen/google/protobuf &&\
			mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
			rm -rf vendor.protogen/protobuf ;\
		fi

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54322 dbname=item-service user=item-service-user password=item-service-password sslmode=disable"

.PHONY: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v