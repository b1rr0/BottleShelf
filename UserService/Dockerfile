FROM golang:1.21
WORKDIR /app
COPY . /app
RUN go build -o app ./cmd
CMD ["./app"]