FROM golang:1.14-alpine

WORKDIR /otp/code/
ADD ./ /otp/code/

RUN apk update && apk upgrade && apk add --no-cache git

RUN go mod download

RUN go build -o bin/meetup cmd/gomeetup/main.go

ENTRYPOINT ["bin/meetup"]
