# Use a minimal Alpine Linux image
FROM alpine:latest

# Create a directory to hold your Go binary
WORKDIR /app

# Copy your built Go binary into the container
COPY ./main /app/main

# Make the binary executable (if needed)
RUN chmod +x /app/main

# Expose the port your application will listen on
EXPOSE 8080

# Define the command to run the application when the container starts
CMD ["/app/main"]
