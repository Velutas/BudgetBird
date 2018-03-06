FROM nimmis/golang:1.9

COPY * /root/

RUN go get -u github.com/sirupsen/logrus
RUN go get -u github.com/go-chi/chi



RUN go build
# RUN go run main.go
