# Start from the official golang image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY ./backend ./backend
COPY ./web ./web

WORKDIR /app/backend

# Build the Go app
RUN go build -o main .

# Expose port 5000 and 80 to the outside world
EXPOSE 5000
EXPOSE 80

# Command to run the executable
CMD ["./main"]

