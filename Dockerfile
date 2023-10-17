FROM golang:1.21-alpine

WORKDIR /usr/src/app

COPY .env .
COPY app .

EXPOSE 8080
CMD [ "./app" ]