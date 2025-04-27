PROTO_SRC=api/proto
PROTO_OUT=internal/protogen

.PHONY: proto-gen-all proto-gen

proto-gen-all: proto-gen

proto-gen:
	protoc --proto_path=$(PROTO_SRC) \
					--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
					--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
					$(PROTO_SRC)/v1/*.proto

