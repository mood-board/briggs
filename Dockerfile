FROM golang:alpine as builder

RUN apk update  && apk add ca-certificates
RUN adduser -D -g '' appuser

COPY . $GOPATH/src/github/ofonimefrancis/brigg/
WORKDIR $GOPATH/src/github/ofonimefrancis/brigg/

RUN go get -d -v 

RUN go build -o mainApp .


#Build a small image from scratch 
FROM scratch 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd


COPY --from=builder mainApp mainApp
USER appuser
ENTRYPOINT ["mainApp"]