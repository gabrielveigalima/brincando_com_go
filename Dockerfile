FROM golang

COPY . /

WORKDIR /

CMD go run hello.go