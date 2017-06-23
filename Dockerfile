FROM golang

ADD . /app
COPY config /config
WORKDIR /app

# install node and node-sass
RUN apt-get update -qq && apt-get install -qq -y \
	nodejs-legacy \
	npm
RUN npm install -g node-sass

# build backend
ENV GOPATH=/app
RUN go get github.com/DeKugelschieber/go-session \
    github.com/DeKugelschieber/go-util \
    github.com/DeKugelschieber/go-resp \
    github.com/go-sql-driver/mysql
RUN go build -ldflags "-s -w" src/main/main.go

# build frontend
RUN node-sass --output /app/public/css /app/public/scss/main.scss
RUN cd /app/public && ./minvue -config=minify.json

# configure and run
ENV ACWEB_HOST=:8080
ENV ACWEB_LOGDIR=/logs
ENV ACWEB_CONFIG_DIR=/config
ENV ACWEB_TLS_PRIVATE_KEY=
ENV ACWEB_TLS_CERT=
ENV ACWEB_DB_USER=root
ENV ACWEB_DB_PASSWORD=
ENV ACWEB_DB_HOST=
ENV ACWEB_DB=acweb

# expose Assetto Corsa folder
VOLUME ["/ac", "/logs", "/config"]

CMD ["/app/main"]
