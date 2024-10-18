# Arguments for use in 'FROM' directive.
ARG ALPINE_VER
ARG GO_VER

# Build stage.
FROM golang:${GO_VER}-alpine${ALPINE_VER} as builder

# Declaration of an argument to be used within a build stage.
#ARG ALPINE_VER
ARG VERSION

RUN apk update && \
    apk upgrade --no-cache && \
    apk add --no-cache \
        npm

WORKDIR /repo
COPY . .

RUN cd web && \
    npm install && \
    npm run build && \
    cd ..

RUN GOOS=linux go build -tags musl -o backend ./cmd/backend

# Build an application image.
FROM alpine:${ALPINE_VER}

# Declaration of an argument to be used within an application stage.
ARG ALPINE_VER
ARG WB_MIRROR

RUN apk update && \
    apk upgrade --no-cache

WORKDIR /app
COPY --from=builder /repo/backend /app/backend

CMD ["/app/backend"]
