FROM golang:1.21.0


WORKDIR /app/

COPY main.go /app/

EXPOSE 3000

CMD [ "go","run","main.go" ]
