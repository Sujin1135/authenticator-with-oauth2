FROM golang:1.19.3-alpine3.15 as builder

LABEL maintainer="Chris <sattlub123@gmail.com>"

WORKDIR /usr/src/app
COPY . .

ENV GO111MODULE=on

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o bin/main cmd/main.go

FROM alpine

COPY --from=builder /usr/src/app/bin/main ./main

EXPOSE 8080

ENTRYPOINT ["./main"]
