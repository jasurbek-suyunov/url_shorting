# workspace (GOPATH) configured at /go
FROM golang:1.19.1-alpine3.16 AS builder

RUN mkdir -p /app
WORKDIR /app


COPY .env ./
COPY go.mod ./

# Copy the local package files to the container's workspace.
COPY . ./

RUN apk add --no-cache tzdata
RUN apk update
RUN apk add make

# installing depends and build
RUN export CGO_ENABLED=0
RUN export GOOS=linux
RUN go build -o ./ /app/cmd/main.go
RUN mv main /

FROM alpine

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /app/.env /
COPY --from=builder /main /

RUN apk add --no-cache tzdata
RUN apk update
RUN apk add make

EXPOSE 8000

CMD ["/main"]
