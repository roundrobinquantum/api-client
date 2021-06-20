FROM golang:1.15-alpine AS builder
RUN mkdir -p /go/src/
RUN CGO_ENABLED=0
RUN GOOS=linux

ENV GOPATH /go
WORKDIR /go/src/

ADD . /go/src/

RUN go build

FROM alpine as runtime

RUN mkdir -p /app

COPY --from=builder /go/src/api-client /app/

RUN chmod +x /app/api-client
WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["/app/api-client"]
