FROM golang

ADD . /app
WORKDIR /app

ENV GOPATH=/app

RUN go get github.com/DeKugelschieber/go-session
RUN go get github.com/DeKugelschieber/go-util
RUN go get github.com/DeKugelschieber/go-resp
RUN go get github.com/go-sql-driver/mysql
RUN go build -ldflags "-s -w" src/main/main.go

ENV ACWEB_HOST=0.0.0.0:8080
ENV ACWEB_LOGDIR=log
ENV ACWEB_TLS_PRIVATE_KEY=
ENV ACWEB_TLS_CERT=
ENV ACWEB_DB_USER=root
ENV ACWEB_DB_PASSWORD=
ENV ACWEB_DB_HOST=
ENV ACWEB_DB=acweb

CMD ["/app/main"]
