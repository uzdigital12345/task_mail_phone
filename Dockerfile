# workspace (GOPATH) configured at /go
FROM golang:1.14.2 as builder


#
RUN mkdir -p $GOPATH/src/github.com/uzdigital12345/task_mail_phone
WORKDIR $GOPATH/src/github.com/uzdigital12345/task_mail_phone

# Copy the local package files to the container's workspace.
COPY . .


RUN     go build cmd/main.go


ENTRYPOINT ["./main"]



