# Start from the official Golang base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code to the working directory
COPY main.go go.mod ./

# Build the Go application
RUN go build -o start .

# Set the command to run the application
CMD ["./start"]
