FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
WORKDIR /home/www/messageboard

COPY . .

RUN go build -o app .

WORKDIR /dist

RUN cp /home/www/messageboard/app .

RUN mkdir src .

RUN cp -r /home/www/messageboard/statics ./statics/
RUN cp -r /home/www/messageboard/templates ./templates/

EXPOSE 8080

CMD ["/dist/app"]