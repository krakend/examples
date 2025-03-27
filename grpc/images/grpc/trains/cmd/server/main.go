package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	trainspb "github.com/krakend/examples/grpc/images/grpc/genlibs/trains"
)

func main() {
	fes := NewTrainsEchoServer()

	s := grpc.NewServer(
		grpc.Creds(loadCredentials()),
		grpc.UnaryInterceptor(checkClientCert),
	)
	trainspb.RegisterTrainsServer(s, fes)

	// TODO: select the listen port
	fmt.Printf("binding to :4243\n")
	ls, err := net.Listen("tcp", ":4243")
	if err != nil {
		fmt.Printf("cannot bind to port: %s\n", err.Error())
		return
	}

	fmt.Printf("running echo server ...\n")
	if err := s.Serve(ls); err != nil {
		fmt.Printf("failed to start server: %s\n", err.Error())
	}
}

func checkClientCert(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// get client tls info
	if p, ok := peer.FromContext(ctx); ok {
		if mtls, ok := p.AuthInfo.(credentials.TLSInfo); ok {
			for _, item := range mtls.State.PeerCertificates {
				fmt.Printf("request certificate subject: %#v", item.Subject)
			}
		}
	}
	return handler(ctx, req)
}

func loadCredentials() credentials.TransportCredentials {
	certificate, err := tls.LoadX509KeyPair("certs/server.cert", "certs/server.key")
	if err != nil {
		panic("failed to load server certification: " + err.Error())
	}

	data, err := ioutil.ReadFile("certs/ca.crt")
	if err != nil {
		panic("failed to load CA file: " + err.Error())
	}

	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(data) {
		panic("can't add ca cert")
	}

	// Uncomment the lines below to enable mTLS:
	tlsConfig := &tls.Config{
		// ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		// ClientCAs:    capool,
	}
	return credentials.NewTLS(tlsConfig)
}
