package main

import (
	"context"
	"fmt"
	"github.com/xiaotuancai/grpc"
	"github.com/xiaotuancai/grpc/credentials"
	"github.com/xiaotuancai/grpc/test/hello"
	"net"
)

type CommServer struct {
}

func (comm *CommServer)Speak(ctx context.Context, content *hello.Content) (*hello.Content, error) {
	fmt.Println("receive message :", content.Detail)
	return &hello.Content{Detail: "i am server"}, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:34516")
	if err != nil {
		panic(err)
	}

	// load tls cert pair  single cert mode
	creds, err := credentials.NewServerTLSFromFile("/opt/gopath/src/github.com/xiaotuancai/grpc/test/single-cert/server.crt",
		"/opt/gopath/src/github.com/xiaotuancai/grpc/test/single-cert/server.key")

	// load tls cert pair double cert mode
	//creds, err := credentials.NewServerTLSFromFileDouble(
	//	"/opt/gopath/src/github.com/xiaotuancai/grpc/test/double-cert/server_sign.crt",
	//	"/opt/gopath/src/github.com/xiaotuancai/grpc/test/double-cert/server_sign.key",
	//	"/opt/gopath/src/github.com/xiaotuancai/grpc/test/double-cert/server_cipher.crt",
	//	"/opt/gopath/src/github.com/xiaotuancai/grpc/test/double-cert/server_cipher.key")
	if err != nil {
		panic(err)
	}

	grpcOptions := []grpc.ServerOption{grpc.Creds(creds)}
	gprcServer := grpc.NewServer(grpcOptions...)
	hello.RegisterCommunicateServer(gprcServer, &CommServer{})
	fmt.Println("beginning to serve ...")
	if err = gprcServer.Serve(l); err != nil {
		panic(err)
	}
}
