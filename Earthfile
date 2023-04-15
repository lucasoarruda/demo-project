VERSION 0.7
all:
    BUILD --platform=linux/amd64 --platform=linux/arm64 +docker
amd64:
    BUILD --platform=linux/amd64 +docker
deps:
    FROM golang:alpine
    WORKDIR /build
    COPY . ./
    RUN apk add --no-cache git
    WORKDIR /build/golang
    RUN go mod tidy
    RUN go mod download
    RUN go get -u github.com/swaggo/swag/cmd/swag
    RUN go install github.com/swaggo/swag/cmd/swag
    RUN swag init -g ../../cmd/web/routes.go -o ./docs -d ./internal/handlers

compile:
    FROM +deps
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG VARIANT
    RUN GOARM=${VARIANT#v} CGO_ENABLED=0 go build \
        --ldflags "-X 'main.Version=v0.0.1' -X 'main.BuildTime=$(date "+%H:%M:%S--%d/%m/%Y")' -X 'main.GitCommit=$(git rev-parse --short HEAD)'" \
        -installsuffix 'static' \
        -o compile/demo-project cmd/web/*.go
    SAVE ARTIFACT compile/demo-project /demo-project AS LOCAL compile/demo-project

docker:
    ARG EARTHLY_TARGET_TAG_DOCKER
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM gcr.io/distroless/static
    #FROM --platform=$TARGETPLATFORM golang:alpine
    WORKDIR /
    COPY \
        --platform=linux/amd64 \
        (+compile/demo-project --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) /demo-project
    ENV GIN_MODE=release
    CMD ["/demo-project"]
    SAVE IMAGE --push ghcr.io/lucasoarruda/demo-project:$EARTHLY_TARGET_TAG_DOCKER