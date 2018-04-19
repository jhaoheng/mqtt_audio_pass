FROM        golang:1.9.4-alpine

# install sound dep
RUN         apk update && apk upgrade && \
            apk add --no-cache bash git openssh && \
            apk add --no-cache gcc musl-dev && \
            apk add alsa-lib-dev

RUN         go get github.com/hajimehoshi/oto
RUN         go get github.com/hajimehoshi/go-mp3
