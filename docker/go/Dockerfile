FROM golang:1.12
ENV GO111MODULE=on

ARG GIT_AUTHOR
ARG GIT_REPOSITORY
WORKDIR /go/src/github.com/${GIT_AUTHOR}/${GIT_REPOSITORY}

RUN go get -u github.com/tockins/realize

CMD ["realize", "start", "--run", "--no-config"]
