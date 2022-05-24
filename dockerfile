FROM ubuntu:latest

RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

RUN apt-get update && apt-get install -y sudo vim curl wget jq git \
    && rm -rf /var/lib/apt/lists/*

ENV PATH="/usr/local/go/bin:/root/go/bin:$PATH"
RUN wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.2.linux-amd64.tar.gz && rm go1.18.2.linux-amd64.tar.gz && \
    go env -w GOBIN=$HOME/go/bin

RUN PREFIX="/usr/local" && \
    VERSION="1.4.0" && \
    curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m).tar.gz" | \
    tar -xvzf - -C "${PREFIX}" --strip-components 1

RUN go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
    go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@v1.4.1-0.20220401124828-90e4c3283529 && \
    go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@v1.0.0-alpha7

CMD ["/bin/bash"]