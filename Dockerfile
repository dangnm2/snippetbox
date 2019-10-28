FROM golang:1.13-alpine

RUN apk --no-cache add make gcc

WORKDIR /app

COPY . .

RUN make build

RUN make test

CMD ["./bin/snippetbox"]
