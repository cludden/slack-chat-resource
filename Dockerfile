FROM alpine

RUN apk update
RUN apk add ca-certificates

COPY check /opt/resource/check
COPY in /opt/resource/in
COPY out /opt/resource/out