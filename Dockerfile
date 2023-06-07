FROM golang:1.20-buster

# Set the working directory
WORKDIR /go/src/chainlink

# Copy the Go module files
COPY chainlink/go-base-executable/go.mod chainlink/go-base-executable/go.sum ./go-base-executable/

# Download dependencies
RUN cd ./go-base-executable && go mod download

# Copy the chainlink client code
COPY go-base-executable/src/chainlink_client.go .


# Build the chainlink client
RUN go build -o mayhemit

# Set the entrypoint
ENTRYPOINT ["./mayhemit"]
