FROM golang:1.16.6-alpine3.14 AS build

ARG COMMIT_ID=missing
LABEL app=k8s-secret-clone vendor=kubernetes version=0.0.1 maintainer=matheus-cal \
      git-revision=$COMMIT_ID \
      documentation=https://github.com/matheus-cal/k8s-secret-cloner


WORKDIR /go/src
RUN apk -U upgrade --no-cache && apk add -U --no-cache \ 
    tree=1.8.0-r0 git=2.32.0-r0 make=4.3-r0 build-base=0.5-r2
COPY src/. ./

ENV GIT_TERMINAL_PROMPT=1 

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o post-install

FROM alpine:3.14 AS deploy

LABEL app=k8s-secret-clone vendor=kubernetes version=0.0.1 maintainer=matheus-cal \
      git-revision=$COMMIT_ID \
      documentation=https://github.com/matheus-cal/k8s-secret-cloner

USER root
RUN apk -U upgrade --no-cache
COPY --from=build /go/src/post-install /usr/local/bin/post-install
USER nobody
CMD ["/usr/local/bin/post-install"]
