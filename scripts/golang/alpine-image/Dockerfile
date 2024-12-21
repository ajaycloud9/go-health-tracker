# Use a minimal Alpine base image
FROM alpine:3.20

# Update and install required packages
RUN apk update -v && \
    apk add --no-cache \
    openssl \
    ca-certificates \
    bash

# Set /bin/sh as the default command (which is already part of Alpine)
CMD ["/bin/sh"]
