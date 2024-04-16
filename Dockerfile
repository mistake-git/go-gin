FROM golang:1.18.1-alpine

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .

# モジュールのダウンロード
RUN go mod download

# ソースコードをコピー
COPY src/ .

# アプリケーションをビルドして実行
CMD ["go", "run", "main.go"]
