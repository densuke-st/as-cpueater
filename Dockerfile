# densukest/cpueater

FROM golang:alpine as builder
WORKDIR /work
COPY loop.go .
RUN go build -o /loop loop.go

FROM alpine
COPY --from=builder /loop .
ENTRYPOINT ["/loop"]
CMD ["/stop"]

