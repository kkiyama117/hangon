FROM golang:latest
MAINTAINER kkiyama117 <k.kiyama117@gmail.com>
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
    apt-utils \
    locales \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
ENV LANG="ja_JP.UTF-8" \
    LANGUAGE="ja_JP:ja" \
    LC_ALL="ja_JP.UTF-8" \
    TZ="JST-9" \
    GO111MODULE=on

# WORKDIR -> $GOPATH
COPY . ./src
WORKDIR ./src
RUN go mod tidy && \
    go build -o pumpkin
CMD ["pumpkin"]
