FROM golang:1.17-alpine AS builder

ENV BIN_FILE /opt/gatekeeper
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} .

FROM alpine:latest
LABEL maintainer="Mayshev Alex <alex.mayshev.86@gmail.com>"

ENV BIN_FILE /opt/gatekeeper
COPY --from=builder ${BIN_FILE} ${BIN_FILE}

EXPOSE 50051

CMD ${BIN_FILE} api