# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the source code into the container
COPY . .

#Initialize go.mod file
RUN go mod init todo-app

# Build the Go application
RUN go build -o todo-app .

# Expose the port the application will run on
EXPOSE 8080

# Define the command to run the application
CMD ["./todo-app"]