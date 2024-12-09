FROM golang:alpine

RUN go install github.com/air-verse/air@latest

RUN mkdir -p /var/www/tolma.ch

COPY . /var/www/tolma.ch

WORKDIR /var/www/tolma.ch

EXPOSE 8000

CMD air
