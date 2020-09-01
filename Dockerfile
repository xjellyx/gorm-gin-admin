FROM  golang:latest

MAINTAINER olongfen

WORKDIR /app

RUN export CONF_DIR=$(pwd)

ADD ./conf/ /app/conf

ADD ./main /app

EXPOSE 8050
EXPOSE 9050

ENTRYPOINT ["./main"]