ARG go_base_image=golang:1.24-alpine

FROM ${go_base_image} AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/contract/deploy.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/data/pumper.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/query/validator.go

FROM alpine

RUN apk add --no-cache tzdata
ENV TZ=Asia/Kolkata

COPY --from=builder /build/deploy /app/deploy
COPY --from=builder /build/pumper /app/pumper
COPY --from=builder /build/validator /app/validator
