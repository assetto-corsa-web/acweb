FROM golang

ADD . /go/src/github.com/assetto-corsa-web/acweb/
COPY configs /config
WORKDIR /go/src/github.com/assetto-corsa-web/acweb/

# install node
RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_8.x -o nodesource_setup.sh && bash nodesource_setup.sh
RUN apt-get install -y nodejs

# build backend
ENV GOPATH=/go
RUN go build -ldflags "-s -w" main.go

# build frontend
RUN cd /go/src/github.com/assetto-corsa-web/acweb/public && npm install && npm run build

# configure and run
ENV ACWEB_HOST=:8080
ENV ACWEB_LOGDIR=/logs
ENV ACWEB_LOGLEVEL=info
ENV ACWEB_INSTANCE_LOGDIR=/instance_logs
ENV ACWEB_CONFIG_DIR=/config
ENV ACWEB_TLS_PRIVATE_KEY=
ENV ACWEB_TLS_CERT=
ENV ACWEB_DB_USER=root
ENV ACWEB_DB_PASSWORD=
ENV ACWEB_DB_HOST=
ENV ACWEB_DB=acweb
ENV ACWEB_DB_PORT=5432
ENV ACWEB_DB_SSLMODE=disable
ENV ACWEB_DB_SSL_CERT=
ENV ACWEB_DB_SSL_KEY=
ENV ACWEB_DB_ROOT_CERT=

# expose Assetto Corsa folder
VOLUME ["/ac", "/logs", "/instance_logs", "/config"]

CMD ["/go/src/github.com/assetto-corsa-web/acweb/main"]
