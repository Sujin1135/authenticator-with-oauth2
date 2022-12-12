### Bulder
FROM golang:1.19.4-alpine3.16 as builder
RUN apk update
RUN apk add git
RUN apk add ca-certificates;

WORKDIR /usr/src/app
COPY . .

ENV GO111MODULE=on

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o bin/main cmd/main.go

### Executable Image
FROM alpine

COPY --from=builder /usr/src/app/bin/main ./main
COPY --from=builder /usr/src/app/.env ./

EXPOSE 8080

ENTRYPOINT ["./main"]
