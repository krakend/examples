package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	flightspb "github.com/krakend/examples/grpc/images/grpc/genlibs/flights"
	libpb "github.com/krakend/examples/grpc/images/grpc/genlibs/lib"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type FlightsEchoServer struct {
	flightspb.UnimplementedFlightsServer
	rnd *rand.Rand
}

func NewFlightsEchoServer() *FlightsEchoServer {
	rndSrc := rand.NewSource(0)
	rnd := rand.New(rndSrc)
	return &FlightsEchoServer{
		rnd: rnd,
	}
}

func prettyPrint(title string, i interface{}) {
	bytesOut, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Printf("cannot display %s\n", err.Error())
		return
	}
	fmt.Printf("* -> [ %s ]:\n%s\n\n", title, string(bytesOut))
}

func (s *FlightsEchoServer) FindFlight(ctx context.Context,
	req *flightspb.FindFlightRequest) (*flightspb.FindFlightResponse, error) {

	tm := time.Now()
	fmt.Printf("\n-[FindFlight @ %v]-----\n", tm)
	prettyPrint("received", req)
	resp := s.generateResponse(req)
	prettyPrint("sending", resp)
	fmt.Printf("\n=[FindFlight @ %v]=====\n", tm)
	return resp, nil
}

func (s *FlightsEchoServer) BookFlight(ctx context.Context,
	req *flightspb.BookFlightRequest) (*flightspb.BookFlightResponse, error) {

	tm := time.Now()
	fmt.Printf("\n-[BookFlight @ %v]-----\n", tm)
	prettyPrint("received", req)
	resp := &flightspb.BookFlightResponse{
		ConfirmationId: fmt.Sprintf("%v", tm),
	}
	prettyPrint("sending", resp)
	fmt.Printf("\n=[BookFlight @ %v]=====\n", tm)
	return resp, nil
}

func (s *FlightsEchoServer) generateResponse(req *flightspb.FindFlightRequest) *flightspb.FindFlightResponse {
	page := s.generateResponsePage(req)
	offset, _ := strconv.ParseInt(page.Cursor, 10, 32)
	if offset > 0 {
		offset -= 1
	}
	return &flightspb.FindFlightResponse{
		Page:    page,
		Flights: s.generateResponseFlights(req, int(page.Size), int(page.Size)*int(offset)),
	}
}

func (s *FlightsEchoServer) generateResponsePage(req *flightspb.FindFlightRequest) *libpb.Page {
	pageCur := req.GetPage().GetCursor()
	pageSize := req.GetPage().GetSize()

	if pageSize < 1 {
		pageSize = 2
	}
	if int(pageSize) > len(airports) {
		pageSize = int32(len(airports))
	}

	if len(pageCur) == 0 {
		pageCur = "1"
	} else {
		if i, err := strconv.ParseInt(pageCur, 10, 32); err != nil {
			pageCur = fmt.Sprintf("%d", i+1)
		} else {
			pageCur = "1"
		}
	}

	return &libpb.Page{
		Size:   pageSize,
		Cursor: pageCur,
	}
}

func (s *FlightsEchoServer) generateResponseFlights(req *flightspb.FindFlightRequest, numFlights, offset int) []*flightspb.FlightInfo {
	flights := make([]*flightspb.FlightInfo, numFlights)
	for i := 0; i < numFlights; i++ {
		flights[i] = s.generateResponseFlight(req, i)
	}
	return flights
}

func (s *FlightsEchoServer) generateResponseFlight(req *flightspb.FindFlightRequest, idx int) *flightspb.FlightInfo {
	return &flightspb.FlightInfo{
		Origin:      s.generateLocation(req, idx),
		Destination: s.generateLocation(req, (idx+1)*(idx+2)),
		Departure:   timestamppb.New(time.Now()),
		Arrival:     timestamppb.New(time.Date(2023, 11, 33, 14, 49, 51, 42, time.UTC)),
	}
}

func (s *FlightsEchoServer) generateLocation(req *flightspb.FindFlightRequest, idx int) *libpb.Location {
	if idx < 0 {
		idx = 0
	}
	idx = idx % len(airports)
	a := airports[idx]
	return &a
}
