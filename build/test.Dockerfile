FROM golang:1.17

ENV CODE_DIR /opt/integration_tests

RUN mkdir -p ${CODE_DIR}
WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .