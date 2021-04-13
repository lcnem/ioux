# Simple usage with a mounted data directory:
# > docker build -t ioux .
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ioux:/root/.ioux ioux iouxd init
# > docker run -it -p 26656:26656 -p 26657:26657 -v ~/.ioux:/root/.ioux ioux iouxd start
FROM golang:1.16-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/lcnem/ioux

# Add source files
COPY . .

RUN go version

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES
RUN make install

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates

WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/iouxd /usr/bin/iouxd

# Run iouxd by default, omit entrypoint to ease using container with iouxcli
CMD ["iouxd"]