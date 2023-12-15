```sh
# Initializes and writes a new buf.yaml configuration file.
buf mod init

# Generate code with protoc plugins
buf generate proto

# Run linting on Protobuf files
buf lint proto

# Build Protobuf files into a Buf image
# Buf image is a binary representation of a compiled Protobuf schema
buf build
```