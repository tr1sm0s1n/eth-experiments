ARG go_base_image=golang:1.25.0-alpine

FROM ${go_base_image} AS builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/transactor/batch_transactor.go

FROM scratch AS runner

COPY ./test.csv .
COPY --from=builder /build/batch_transactor .
