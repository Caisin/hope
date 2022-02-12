# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17 AS build

COPY ../ /hope
WORKDIR /
ENV GOPROXY=https://goproxy.cn
RUN  cd /hope && go mod download && \
    cd /hope/apps/admin/cmd/admin/ && go build -o admin && cp admin / && \
    cd /hope/apps/novel/cmd/novel/ && go build -o novel && cp novel / && \
    cd /hope/apps/param/cmd/param/ && go build -o param && cp param /


##
## Deploy
##
FROM debian:stable-slim

WORKDIR /app

COPY --from=build /admin /novel /param /app/

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

#USER nonroot:nonroot
#CMD ["./admin", "-conf", "/data/conf"]
#ENTRYPOINT ["/docker-gs-ping"]