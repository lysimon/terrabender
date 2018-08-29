FROM golang:1.11.0 as builder
WORKDIR /builder/

RUN go get -d -v golang.org/x/net/html &&\
    go get -d -v github.com/robfig/cron &&\
    go get -d -v github.com/aws/aws-sdk-go &&\
    go get github.com/gorilla/mux &&\
    go get -d -v gopkg.in/yaml.v2 &&\
    go get -d -v github.com/google/go-github/github

COPY ./cmd    /go/src/github.com/lysimon/terrabender/cmd
COPY ./pkg    /go/src/github.com/lysimon/terrabender/pkg

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app /go/src/github.com/lysimon/terrabender/cmd/terrabender.go

FROM alpine:latest
# Install latest terraform
RUN apk update && apk upgrade && \
    apk --no-cache add ca-certificates bash git openssh terraform wget curl
# Install latest terragrunt

# Install terragrunt
# RUN apt-get install wget
WORKDIR /root/

COPY --from=builder /builder/app .
CMD ["./app"]
