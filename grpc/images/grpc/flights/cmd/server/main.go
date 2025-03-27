package main

import (
	"fmt"
	"net"

	flightspb "github.com/krakend/examples/grpc/images/grpc/genlibs/flights"
	"google.golang.org/grpc"
)

func main() {
	fes := NewFlightsEchoServer()

	s := grpc.NewServer()
	flightspb.RegisterFlightsServer(s, fes)

	// TODO: select the listen port
	fmt.Printf("binding to :4242")
	ls, err := net.Listen("tcp", ":4242")
	if err != nil {
		fmt.Printf("cannot bind to port: %s\n", err.Error())
		return
	}

	fmt.Printf("running echo server ...\n")
	if err := s.Serve(ls); err != nil {
		fmt.Printf("failed to start server", err.Error())
	}
}
