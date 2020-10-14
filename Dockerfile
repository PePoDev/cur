# build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
ADD . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -i -a -o /go/bin/cur .

# final stage
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/cur /
CMD ["/cur", "report", ""]