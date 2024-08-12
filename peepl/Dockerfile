# Use the official Golang image to build the app
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Create the /etc/peepl directory
RUN mkdir -p /etc/peepl

# Copy the local config.yaml file into the container
COPY config.yaml /etc/peepl/config.yaml

# Copy migrations directory into the container
COPY db/migration ./db/migration

# Copy the wait-for-it.sh script into the container
COPY wait-for-it.sh /root/wait-for-it.sh

# Make the wait-for-it.sh CRLF to LF
RUN sed -i 's/\r$//' wait-for-it.sh

# Give execute permission to the wait-for-it.sh script
RUN chmod +x /root/wait-for-it.sh

# Command to run the executable
CMD ["./wait-for-it.sh", "db:5432", "--timeout=60", "--strict", "--", "./main"]
