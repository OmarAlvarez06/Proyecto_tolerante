FROM golang:1.17.3-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/Pruebas /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./Pruebas"]