FROM alpine:latest 

ENV PROJECT_NAME=rest-service

RUN mkdir -p /usr/bin/testdata
COPY bin/service-test /usr/bin/service-test