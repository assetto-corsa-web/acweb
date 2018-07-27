FROM golang

ADD . /app
COPY configs /config
WORKDIR /app

# install sass
RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y \
	build-essential \
	ruby-dev
RUN gem install sass

# build backend
ENV GOPATH=/app
RUN go build -ldflags "-s -w" src/main/main.go

# build frontend
RUN cd /app/public && mkdir ./css && sass ./scss/main.scss ./css/main.css
RUN cd /app/public && ./minvue -config=minify.json

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

CMD ["/app/main"]
