FROM golang:1.19

WORKDIR /go/src/app
COPY ./tracking_impression .

RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
ENV PATH="${PATH}:${GOPATH}/bin"

CMD ["./tracking"]
