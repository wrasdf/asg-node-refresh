# Build the manager binary
FROM golang:1.15 as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o ./bin/asg-node-refresh main.go

# Release binary
FROM alpine:3.13 as release
WORKDIR /app

COPY --from=Builder /app/bin/asg-node-refresh .

EXPOSE 8085
ENTRYPOINT ["/app/asg-node-refresh"]
