FROM golang:alpine as builder
WORKDIR /go/src/app
COPY main.go .
RUN go build .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/app/app /app

CMD ["/app"]

