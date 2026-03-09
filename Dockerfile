# Use official Go image as base
FROM golang:1.25.6

# Set working directory inside container
WORKDIR /app

# Copy source code
COPY . .

# Build the binary
RUN go build -o payment-sim

# Run the binary when container starts
CMD ["./payment-sim"]
