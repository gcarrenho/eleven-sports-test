FROM golang:alpine3.15 as build

WORKDIR /go/src/app

COPY . .

#RUN go mod init

RUN apk add git

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go build -o main ./src/cmd

# Running containers as non-root user
RUN chown -R 1002 /go/src   
USER 1002

ENTRYPOINT ["/go/src/app/main"]