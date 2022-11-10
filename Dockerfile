FROM golang:1.19-alpine
WORKDIR /app
COPY * ./
RUN go env -w GO111MODULE=off
RUN go build -o /fizzbuzz
CMD [ "/fizzbuzz" ]

