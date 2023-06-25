FROM golang:alpine AS builder

LABEL stage=gobuilder

WORKDIR /myapp/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /myapp/release
COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app ./cmd/server/main.go

FROM alpine:latest as prod
WORKDIR /myapp
COPY --from=builder /myapp/release/. /myapp

CMD ["/myapp/app"]