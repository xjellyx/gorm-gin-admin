FROM  golang:latest

MAINTAINER olongfen

WORKDIR /app

RUN export CONF_DIR=$(pwd)

ADD ./conf/ /app/conf

ADD ./main /app

EXPOSE 8060
EXPOSE 9060

ENTRYPOINT ["./main"]