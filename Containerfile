FROM registry.access.redhat.com/ubi9/ubi-micro:9.6-1752751762

WORKDIR /go/src/app
COPY ./oaimg-service .

USER nobody
EXPOSE 8080
CMD ["./oaimg-service"]
