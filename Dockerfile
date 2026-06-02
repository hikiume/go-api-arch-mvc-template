# 軽量なAlpine LinuxベースのGoイメージを使用
FROM golang:1.24-alpine

# oapi-codegenのインストールに必要なツールとoapi-codegen本体をインストール
RUN apk update && apk add --no-cache git \
    && go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# コンテナが起動した時のデフォルトのコマンド（シェルを起動）
CMD ["sh"]
