FROM golang:1.23-alpine AS builder

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

WORKDIR /app/

COPY --from=builder /opt/gatekeeper .
COPY --from=builder /go/src/.env .

EXPOSE 50051

CMD ["./gatekeeper", "api"]