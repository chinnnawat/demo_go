FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY ./core/go.mod ./
COPY ./core/go.sum ./
RUN go mod download

COPY ./core ./
COPY ./core/.air.toml ./

RUN apk add --no-cache shadow bash libstdc++ gcc g++

RUN adduser --disabled-password --gecos "" vscode
RUN groupadd docker
RUN usermod -aG docker vscode

# แก้ไขสิทธิ์ของโฟลเดอร์ต่างๆ และแน่ใจว่าผู้ใช้ vscode สามารถเข้าถึงได้
USER root
RUN mkdir -p /go/pkg/mod && \
    chown -R vscode:vscode /go/pkg/mod && \
    chown -R vscode:vscode /go/pkg && \
    mkdir -p /home/vscode/.vscode-server && \
    chown -R vscode:vscode /home/vscode/.vscode-server && \
    chown -R vscode:vscode /app && \
    chmod -R 755 /app/models /app/tmp

USER vscode

CMD ["air", "-c", ".air.toml"]
