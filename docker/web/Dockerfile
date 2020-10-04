FROM golang:1.15.2

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV server.port=8080
ENV database.dialect=sqlserver
ENV database.user=sa
ENV database.password=123456Aa
ENV database.port=1433
ENV database.server=db.dev
ENV database.database=BLOG_DEV
CMD ["blog"]