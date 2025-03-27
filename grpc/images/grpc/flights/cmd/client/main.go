package main

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	flightspb "github.com/krakend/examples/grpc/images/grpc/genlibs/flights"
	libpb "github.com/krakend/examples/grpc/images/grpc/genlibs/lib"
	// timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	fmt.Printf("this is a client...\n")
	c := NewFlightsClient("localhost:4242")
	c.FindFlight()
	c.BookFlight()
}

type FlightsClient struct {
	conn flightspb.FlightsClient
}

func NewFlightsClient(addr string) *FlightsClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("CANNOT Dial : %s\n", err.Error())
		return nil
	}
	return &FlightsClient{
		conn: flightspb.NewFlightsClient(conn),
	}
}

func (c *FlightsClient) FindFlight() {
	ctx := context.Background()
	resp, err := c.conn.FindFlight(ctx, &flightspb.FindFlightRequest{
		Page: &libpb.Page{
			Size:   20,
			Cursor: "foo",
		},
	})

	if err != nil {
		fmt.Printf("\ncannot get flights %s\n", err.Error())
		return
	}
	prettyPrint("Flights", resp)
}

func (c *FlightsClient) BookFlight() {
	ctx := context.Background()
	resp, err := c.conn.BookFlight(ctx, &flightspb.BookFlightRequest{
		FlightId: "foobar",
		Passengers: []*flightspb.Passenger{
			&flightspb.Passenger{
				FullName: "Foo",
			},
		},
	})

	if err != nil {
		fmt.Printf("\ncannot book flight %s\n", err.Error())
		return
	}
	prettyPrint("Flights", resp)

}

func prettyPrint(title string, i interface{}) {
	bytesOut, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Printf("cannot display %s\n", err.Error())
		return
	}
	fmt.Printf("\n**[ %s ]**\n%s\n__[%s]__\n", title, string(bytesOut), title)
}
