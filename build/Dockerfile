# alpineは軽量に動作するので今回は採用
FROM golang:1.14.1-alpine

# alpineはGitやVim等のオプションツールは入っていないのでインストール
# 改行しないとエラーになる
RUN apk --update-cache --no-cache add \
    git
