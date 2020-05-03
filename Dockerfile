FROM golang:alpine

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN mkdir -p /go/src/github.com/emre/react-golang-web-app

# Install dep dependency management
RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/emre/react-golang-web-app

WORKDIR /go/src/github.com/emre/react-golang-web-app

# Fetch dependencies.
RUN dep ensure -vendor-only

# Build the binary.
RUN go build
CMD ["./react-golang-web-app"]
