# build the server binary
FROM golang:1.10.0 AS builder
LABEL stage=server-intermediate
WORKDIR /go/src//catalog
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server ./cmd/server

# copy the server binary from builder stage; run the server binary
FROM alpine:latest AS runner
WORKDIR /bin
COPY --from=builder /go/src//catalog/bin/server .
COPY pkg/pb/*.swagger.json tmp/www/swagger
COPY --from=builder /go/src//catalog/db/migrations /db/migrations/
ENTRYPOINT ["server", "--gateway.swaggerFile", "www/catalog.swagger.json"]
