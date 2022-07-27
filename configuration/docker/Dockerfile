# Latest golang as builder
FROM golang:1.18-alpine as builder

# Install build depenedencies
RUN apk --no-cache add git nano

# Set workdir
WORKDIR /root

# Try to get upstream version (default v0.37.0-beta)
ARG UPSTREAM_VER=v0.37.0-beta

# Clone and make TrueBlocks Core
# make -j 5 is a fairly safe number
RUN git clone -b "${UPSTREAM_VER}" --single-branch --progress --depth 1 \
        https://github.com/TrueBlocks/trueblocks-dappnode.git  && \
        cd trueblocks-dappnode/configuration && go build -o server

# Switch to alpine container
FROM alpine:latest

# Install binary dependencies and nice to haves
RUN apk --no-cache add nano

# Copy files from builder
COPY --from=builder /root/trueblocks-dappnode/configuration/server /app/server
COPY --from=builder /root/trueblocks-dappnode/configuration/static /app/static

EXPOSE 80

# Run entrypoint
ENTRYPOINT /app/server
