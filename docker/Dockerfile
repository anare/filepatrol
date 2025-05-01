# Use official Go image as base
FROM golang:1.23


# Set working directory
WORKDIR /app

# Install build dependencies
RUN apt-get update && \
    apt-get install -y build-essential gcc zip git && \
    apt-get clean

# Default entrypoint can be overridden by docker-compose
ENTRYPOINT ["/bin/sh", "-c"]

ARG UID=1000
RUN useradd -u $UID -m builder
USER builder
