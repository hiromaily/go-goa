ARG GOLAGN_VER=1.19.2
ARG ALPINE_VER=3.16

FROM golang:${GOLAGN_VER}-alpine${ALPINE_VER} AS builder
#RUN apk add --no-cache bash jq

WORKDIR /workspace

COPY ./ ./

# First, build code to generate keys
RUN go build -v -o ${GOPATH}/bin/goa-server ./cmd/resume/server/...


# Go Server
FROM alpine:${ALPINE_VER}

WORKDIR /workspace

COPY --from=builder /go/bin/goa-server /usr/bin/goa-server
COPY docs ./
COPY configs/docker.toml ./configs/

EXPOSE 8090

ENTRYPOINT ["goa-server"]