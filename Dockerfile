FROM golang:1.20 AS builder
ENV CGO_ENABLED 0
WORKDIR /go/src/app
ADD . .
RUN go build -o /dellog

FROM busybox
COPY --from=builder /dellog /dellog
CMD ["/dellog"]
