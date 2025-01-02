# ---------------------------------------------------------#
#                   Build image                            #
# ---------------------------------------------------------#
FROM --platform=$BUILDPLATFORM golang:1.23-alpine3.20 AS builder
RUN apk update && apk add --no-cache build-base musl-dev

WORKDIR /app
# RUN git config --global --add safe.directory '/app'

COPY go.mod .
COPY go.sum .


# ---------------------------------------------------------#
#            COPY the source code as the last step         #
# ---------------------------------------------------------#
COPY . .

# ---------------------------------------------------------#
#                           Build the app                  #
# ---------------------------------------------------------#
ARG TARGETOS TARGETARCH

RUN if [ "$TARGETARCH" = "arm64" ]; then \
    wget -P ~ https://musl.cc/aarch64-linux-musl-cross.tgz && \
    tar -xvf ~/aarch64-linux-musl-cross.tgz -C ~ ; \
    fi

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    if [ "$TARGETARCH" = "arm64" ]; then CC=~/aarch64-linux-musl-cross/bin/aarch64-linux-musl-gcc; fi && \
    LDFLAGS="-X github.com/mujhtech/b0/cmd/version.Version=$(git describe --tags) -extldflags '-static'" && \
    CGO_ENABLED=1 \
    GOOS=$TARGETOS GOARCH=$TARGETARCH \
    CC=$CC go build -ldflags="$LDFLAGS" -o ./b0 ./cmd


FROM --platform=$TARGETPLATFORM alpine/git:2.43.0 AS final

# ---------------------------------------------------------#
#                  setup app dir and its content           #
# ---------------------------------------------------------#
WORKDIR /app
VOLUME /data

COPY --from=builder /app/b0 /app/b0

EXPOSE 5555