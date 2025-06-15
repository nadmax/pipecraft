FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY ./src/ ./src/
WORKDIR /app/src
RUN go build -ldflags "-s -w" -o /app/bin/app main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/bin/app /app
USER nonroot:nonroot
ENTRYPOINT ["/app"]
