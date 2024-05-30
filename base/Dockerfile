FROM golang:alpine3.16 AS builder
WORKDIR /code
COPY main.go .
RUN go build -trimpath -o /app main.go

FROM alpine:3
ENV GOTRACEBACK=single
CMD ["./app"]
COPY --from=builder /app .
