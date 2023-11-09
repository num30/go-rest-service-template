FROM alpine:latest 

ENV PROJECT_NAME=rest-service

# you may need some data for test so you can copy it to the container here 
# COPY testdata  /usr/bin/testdata

COPY bin/service-test /usr/bin/service-test
CMD ["/usr/bin/service-test"]