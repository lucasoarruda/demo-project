# syntax=docker/dockerfile:1

# Build stage runs on the native builder platform and cross-compiles, so
# multi-arch builds need no QEMU emulation.
FROM --platform=$BUILDPLATFORM cgr.dev/chainguard/go:latest-dev AS build
USER root
WORKDIR /src
ENV CGO_ENABLED=0 \
    GOMODCACHE=/go/pkg/mod \
    GOCACHE=/root/.cache/go-build

COPY golang/go.mod golang/go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOBIN=/usr/local/bin go install github.com/swaggo/swag/cmd/swag@v1.16.6

COPY golang/ ./
RUN --mount=type=cache,target=/go/pkg/mod \
    swag init -g ../../cmd/web/routes.go -o ./docs -d ./internal/handlers

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT
ARG VERSION=v0.0.4
ARG GIT_COMMIT=unknown
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=$TARGETOS GOARCH=$TARGETARCH GOARM=${TARGETVARIANT#v} \
    go build -trimpath \
      -ldflags "-s -w -X 'main.Version=${VERSION}' -X 'main.BuildTime=$(date '+%H:%M:%S--%d/%m/%Y')' -X 'main.GitCommit=${GIT_COMMIT}'" \
      -o /out/demo-project ./cmd/web

# Wolfi-based distroless runtime; runs as nonroot (UID 65532).
FROM cgr.dev/chainguard/static:latest
COPY --from=build /out/demo-project /demo-project
ENV GIN_MODE=release
USER 65532:65532
EXPOSE 8000
CMD ["/demo-project"]
