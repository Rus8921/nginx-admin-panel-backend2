# Arguments for use in 'FROM' directive.
ARG ALPINE_VER
ARG GO_VER
ARG TARGETARCH
ARG TARGETOS

# Build stage.
FROM golang:${GO_VER}-alpine${ALPINE_VER} AS builder

# Declaration of an argument to be used within a build stage.
ARG VERSION

RUN apk update && \
    apk upgrade --no-cache && \
    apk add --no-cache npm

WORKDIR /repo
COPY . .

# Build the frontend with npm
RUN cd web && \
    npm install && \
    npm run build && \
    cd ..

# Build the Go application for the specified architecture
RUN  go build -tags musl -o backend ./cmd/backend

# Build an application image.
FROM alpine:${ALPINE_VER}

# Declaration of an argument to be used within an application stage.
ARG ALPINE_VER
ARG WB_MIRROR

RUN apk update && \
    apk upgrade --no-cache

WORKDIR /app
COPY --from=builder /repo/backend /app/backend
RUN chmod +x /app/backend

CMD ["/app/backend"]