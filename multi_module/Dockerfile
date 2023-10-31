FROM golang:1.21 AS builder
# ワーキングディレクトリを設定
WORKDIR /app
# ホストのGoモジュールをコンテナにコピー
COPY go.mod .
COPY go.sum .
# モジュールのダウンロード
RUN go mod download

# ビルド引数で指定したモジュールのソースコードをコンテナにコピー
ARG MODULE
COPY ./jobs/${MODULE} ./jobs/${MODULE}
COPY ./shared ./shared
RUN go build -o main ./jobs/${MODULE}

# 実行ステージ
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
# 実行
CMD ["./main"]
