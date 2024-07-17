# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang

ARG OS=linux
ARG ARCH=amd64
FROM golang:1.21.6 as builder

# Set the Current Working Directory inside the container
WORKDIR /app/gateway

# Copy go mod and sum files
COPY ./gateway/go.mod ./gateway/go.sum  /app/gateway/

# Copy go-kit and manager will be removed later
COPY ./go-kit/ /app/go-kit/
COPY ./manager/ /app/manager/
COPY ./sdk/ /app/sdk/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./gateway /app/gateway/

# Build the Go app
#RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -a -installsuffix cgo -o manager .
RUN CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} go build -o gateway ./cmd/main.go 
# Start a new stage from scratch
FROM alpine:latest  

# Add ca-certificates
RUN apk --no-cache add ca-certificates

# Create a dedicated user and group to run the service
# Avoid running as root for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/gateway/gateway .

# Change ownership of the work directory to the appuser
RUN chown -R appuser:appgroup /home/appuser

# Use the created user to run the service
USER appuser

EXPOSE 7777

# Command to run the executable
CMD ["/home/appuser/gateway"]