FROM alpine
# Fixing react version for npm to use this in builder image
RUN apk update && apk add --no-cache --update nodejs=22.11.0-r0
RUN apk add --no-cache --update npm=10.9.1-r0
RUN apk add --no-cache --update libcrypto3=3.3.2-r4
RUN apk add --no-cache --update libssl3=3.3.2-r4
RUN  node --version && npm --version

# Set /bin/sh as the default command (which is already part of Alpine)
CMD ["/bin/sh"]