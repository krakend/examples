package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	// "google.golang.org/grpc/credentials/insecure"

	libpb "github.com/krakend/examples/grpc/images/grpc/genlibs/lib"
	trainspb "github.com/krakend/examples/grpc/images/grpc/genlibs/trains"
)

func main() {
	fmt.Printf("this is a client...\n")
	addr := "localhost:4243"
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(loadCredentials()))
	// grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("CANNOT Dial %s: %s\n", addr, err.Error())
		return
	}
	c := trainspb.NewTrainsClient(conn)

	ctx := context.Background()
	y := int32(2032)
	m := int32(12)
	d := int32(22)
	resp, err := c.FindTrains(ctx, &trainspb.FindTrainRequest{
		Page: &libpb.Page{
			Size:   20,
			Cursor: "foo",
		},
		Origin:      &libpb.Location{},
		Destination: &libpb.Location{},
		Departure: &trainspb.Date{
			Year:  &y,
			Month: &m,
			Day:   &d,
		},
		Arrival: &trainspb.Date{
			Year:  &y,
			Month: &m,
			Day:   &d,
		},
		Classes: []trainspb.Class{},
	})

	if err != nil {
		fmt.Printf("\ncannot get trains %s\n", err.Error())
		return
	}
	prettyPrint("Trains", resp)
}

func prettyPrint(title string, i interface{}) {
	bytesOut, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Printf("cannot display %s\n", err.Error())
		return
	}
	fmt.Printf("\n**[ %s ]**\n%s\n__[%s]__\n", title, string(bytesOut), title)
}

func loadCredentials() credentials.TransportCredentials {

	data, err := ioutil.ReadFile("certs/ca.crt")
	if err != nil {
		panic("failed to load CA file: " + err.Error())
	}
	capool := x509.NewCertPool()
	if !capool.AppendCertsFromPEM(data) {
		panic("can't add ca cert")
	}

	tlsConfig := &tls.Config{
		RootCAs: capool,
	}
	return credentials.NewTLS(tlsConfig)
}
