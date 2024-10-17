FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY ./core/go.mod ./
COPY ./core/go.sum ./
RUN go mod download

COPY ./core ./
COPY ./core/.air.toml ./

RUN apk add --no-cache shadow
RUN apk add --no-cache bash

RUN adduser --disabled-password --gecos "" vscode
RUN groupadd docker
RUN usermod -aG docker vscode

USER root
RUN mkdir -p /home/vscode/.vscode-server && \
    chown -R vscode:vscode /home/vscode/.vscode-server

USER vscode

CMD ["air", "-c", ".air.toml"]
