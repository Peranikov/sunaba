.PHONY: get-tools
get-tools:
	@go get -u \
		google.golang.org/grpc \
		github.com/golang/protobuf/protoc-gen-go

.PHONY: protoc
protoc:
	@protoc -I proto/ --go_out=plugins=grpc:lib proto/*.proto
