FROM golang:alpine

# Set env variables
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Set the working directory inside the container
WORKDIR /app

# Add folder for the build file
RUN mkdir "build"

# Copy the necessary files to the container
COPY . .

# Install CompileDeamon to enable Live Reload
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

# Expose port 8080 for the webserver
EXPOSE 8080

# Command to run the Go application
ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app" -color="true"
