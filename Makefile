.PHONY: get-tools
get-tools:
	@go get -u \
		google.golang.org/grpc \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/99designs/gqlgen

.PHONY: protoc
protoc:
	@protoc -I grpc/proto/ --go_out=plugins=grpc:grpc/lib grpc/proto/*.proto

.PHONY: gqlgen
gqlgen:
	@cd graghql; go run github.com/99designs/gqlgen -v

.PHONY: dep
dep:
	@dep ensure -v

.PHONY: run
run:
	@go run main.go
