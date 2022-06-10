TARGET =

.PHONY: clean build run

build:
	@go build -o ./bin/grpcd ./cmd/grpcd

run:
	@go run ./cmd/$(TARGET)

protoc-gen:
	@./scripts/protoc-build.sh