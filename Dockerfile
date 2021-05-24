FROM golang:1.16 as builder
WORKDIR /app
ADD main.go .
RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
RUN ls .
CMD ["/app/main"]
