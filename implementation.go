package grpctest

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func NewLocalListener() net.Listener {
	if *serve != "" {
		l, err := net.Listen("tcp", *serve)
		if err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on %v: %v", *serve, err))
		}
		return l
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on a port: %v", err))
		}
	}
	return l
}

// When debugging a particular gRPC server-based test,
// this flag lets you run
//	go test -run=BrokenTest -grpctest.serve=127.0.0.1:8000
// to start the broken server so you can interact with it manually.
var serve = flag.String("grpctest.serve", "", "if non-empty, grpctest.NewLocalListener listens on this address and blocks")

// Use to compare two grpc errors bassedon their code,
// ignoring the message
func ErrEqual(e1, e2 error) bool {
	return (e1 == nil && e1 == e2) ||
		(e1 != nil && e2 != nil &&
			grpc.Code(e1) == grpc.Code(e2))
}
