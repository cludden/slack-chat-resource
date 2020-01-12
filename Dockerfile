FROM alpine

RUN apk update
RUN apk add ca-certificates

COPY . /opt/resource