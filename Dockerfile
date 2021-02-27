# Compile stage
FROM golang:1.14.11 AS build-env
RUN mkdir -p /jsoncli
WORKDIR /jsoncli
ADD . /jsoncli
RUN go build search.go
CMD ["./search"]