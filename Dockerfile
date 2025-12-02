FROM golang:1.25 AS builder
WORKDIR /app
RUN apt update
RUN apt install gcc musl-dev sqlite3
COPY go.mod go.sum .
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
ENV CGO_ENABLED=1
RUN go build -ldflags '-linkmode external -extldflags "-static"' -o minimeter ./cmd/minimeter

FROM alpine:3.22
RUN apk add --no-cache sqlite
COPY --from=builder /app /app
ENV DB_PATH=/app/data/
VOLUME ["/app/data"]
EXPOSE 8080
CMD ["/app/minimeter"]
