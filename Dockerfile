FROM golang:1.20-buster as go-target

# Set up the environment
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/chainlink

# Copy the client source code
COPY . .

# Copy the go-base-executable files
COPY ./go-base-executable /go/go-base-executable

# Install Go 1.20
RUN curl -OL https://dl.google.com/go/go1.20.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz && \
    rm go1.20.linux-amd64.tar.gz

# Set the Go 1.20 path
ENV PATH=/usr/local/go/bin:$PATH

# Copy Go module files if required
COPY ./go-base-executable/go.mod /go/go-base-executable/go.sum /go/go-base-executable/

# Set the working directory for building the executable
WORKDIR /go/go-base-executable

# Install Go dependencies if required
RUN go mod download

# Build the go-base-executable
RUN go build -o mayhemit

# Set the entrypoint
ENTRYPOINT ["./mayhemit"]
