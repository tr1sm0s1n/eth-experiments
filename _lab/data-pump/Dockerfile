ARG go_base_image=golang:1.24-alpine

FROM ${go_base_image} AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GO_BUILDFLAGS="-trimpath -ldflags=-s -ldflags=-w"
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build $GO_BUILDFLAGS -o deploy cmd/contract/deploy.go && \
    go build $GO_BUILDFLAGS -o pumper cmd/data/pumper.go && \
    go build $GO_BUILDFLAGS -o injector cmd/api/injector.go && \
    go build $GO_BUILDFLAGS -o validator cmd/query/validator.go && \
    go build $GO_BUILDFLAGS -o server api/server.go

FROM alpine

RUN apk add --no-cache tzdata
ENV TZ=Asia/Kolkata

COPY --from=builder /build/deploy /app/deploy
COPY --from=builder /build/pumper /app/pumper
COPY --from=builder /build/server /app/server
COPY --from=builder /build/injector /app/injector
COPY --from=builder /build/validator /app/validator
