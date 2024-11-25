FROM golang:latest

RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install wkhtmltopdf -y
RUN apt-get install xfonts-75dpi


WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /main main.go

EXPOSE 8082
ENTRYPOINT ["/main"]
