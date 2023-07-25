# Stage 1: Build the Golang binary
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the application source code to the container's working directory
COPY . .

# Build the Golang binary with CGO disabled to create a statically linked binary
RUN CGO_ENABLED=0 go build -o owlery .


# Stage 2: Create a minimal image using Alpine
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Golang binary from the previous stage to the current stage
COPY --from=builder /app/owlery .

# Copy the configuration file from the host to the container
COPY ./configs/application-local.yml /app/configs/application-local.yml

# Set environment variables
ENV CONFIG_PATH /app/configs/application-local.yml
ENV OMETRIA_APIKEY <ometria-api-key>
ENV MAILCHIMP_APIKEY <mailchimp-api-key>

# Run the Golang binary and fetch configurations from the specified file
CMD ["./owlery"]