VERSION 0.8
all:
    BUILD --platform=linux/amd64 --platform=linux/arm64 +docker
amd64:
    BUILD --platform=linux/amd64 +docker
deps:
    # Wolfi-based Chainguard Go image. The free tier only publishes :latest / :latest-dev,
    # so the Go toolchain here tracks the newest release; go.mod's `go` directive sets the floor.
    FROM cgr.dev/chainguard/go:latest-dev
    USER root
    WORKDIR /build
    COPY . ./
    WORKDIR /build/golang
    RUN go mod download
    RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6
    RUN swag init -g ../../cmd/web/routes.go -o ./docs -d ./internal/handlers

compile:
    FROM +deps
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG VARIANT
    RUN GOARM=${VARIANT#v} CGO_ENABLED=0 go build \
        -trimpath \
        --ldflags "-s -w -X 'main.Version=v0.0.3' -X 'main.BuildTime=$(date "+%H:%M:%S--%d/%m/%Y")' -X 'main.GitCommit=$(git rev-parse --short HEAD)'" \
        -o compile/demo-project ./cmd/web
    SAVE ARTIFACT compile/demo-project /demo-project AS LOCAL compile/demo-project

docker:
    ARG EARTHLY_TARGET_TAG_DOCKER
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    # Wolfi-based distroless runtime; runs as nonroot (UID 65532).
    FROM --platform=$TARGETPLATFORM cgr.dev/chainguard/static:latest
    WORKDIR /
    COPY \
        --platform=linux/amd64 \
        (+compile/demo-project --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) /demo-project
    ENV GIN_MODE=release
    USER 65532:65532
    CMD ["/demo-project"]
    SAVE IMAGE --push ghcr.io/lucasoarruda/demo-project:$EARTHLY_TARGET_TAG_DOCKER
