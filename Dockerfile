FROM golang:1.21rc2-bullseye

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the container
COPY go.mod go.sum ./
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the webserver
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]