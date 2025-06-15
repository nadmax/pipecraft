FROM rust:1.87.0-alpine AS builder
WORKDIR /app
RUN apk add --no-cache musl-dev openssl-dev pkgconfig
COPY ./ ./
RUN cargo build --release

FROM alpine:3.22.0
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/target/release/rust /app/app
ENTRYPOINT ["./app"]
