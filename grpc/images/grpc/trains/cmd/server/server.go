package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	libpb "github.com/krakendio/playground-enterprise/images/grpc/genlib/lib"
	trainspb "github.com/krakendio/playground-enterprise/images/grpc/genlib/trains"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type TrainsEchoServer struct {
	trainspb.UnimplementedTrainsServer
	rnd *rand.Rand
}

func NewTrainsEchoServer() *TrainsEchoServer {
	rndSrc := rand.NewSource(0)
	rnd := rand.New(rndSrc)
	return &TrainsEchoServer{
		rnd: rnd,
	}
}

func prettyPrint(title string, i interface{}) {
	bytesOut, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Printf("cannot display %s\n", err.Error())
		return
	}
	fmt.Printf("\n**[ %s ]**\n%s\n\n", title, string(bytesOut))
}

func (s *TrainsEchoServer) FindTrains(ctx context.Context,
	req *trainspb.FindTrainRequest) (*trainspb.FindTrainResponse, error) {

	fmt.Printf("\n-----[%v]-----\n", time.Now())
	prettyPrint("received", req)
	resp := s.generateResponse(req)
	prettyPrint("sending", resp)
	return resp, nil
}

func (TrainsEchoServer) GetTrainClasses(ctx context.Context,
	_ *emptypb.Empty) (*trainspb.TrainClasses, error) {

	mealIncluded := true
	preferentOnboard := false
	return &trainspb.TrainClasses{
		Classes: []trainspb.Class{
			trainspb.Class_REGIONAL,
			trainspb.Class_NATIONAL,
			trainspb.Class_INTERNATIONAL,
		},
		Perks: &trainspb.TrainClasses_Perks{
			MealIncluded:     &mealIncluded,
			PreferentOnboard: &preferentOnboard,
		},
	}, nil
}

func (s *TrainsEchoServer) generateResponse(req *trainspb.FindTrainRequest) *trainspb.FindTrainResponse {
	return &trainspb.FindTrainResponse{
		Page:   s.generateResponsePage(req),
		Trains: s.generateResponseTrains(req),
	}
}

func (s *TrainsEchoServer) generateResponsePage(req *trainspb.FindTrainRequest) *libpb.Page {
	pageCur := req.GetPage().GetCursor()
	if len(pageCur) == 0 {
		pageCur = "1"
	} else {
		if i, err := strconv.ParseInt(pageCur, 10, 32); err != nil {
			pageCur = fmt.Sprintf("%d", i+1)
		} else {
			pageCur = fmt.Sprintf("%s_1")
		}
	}
	pageSize := req.GetPage().GetSize()
	if pageSize == 0 {
		pageSize = 10
	}
	return &libpb.Page{
		Size:   pageSize,
		Cursor: pageCur,
	}
}

func (s *TrainsEchoServer) generateResponseTrains(req *trainspb.FindTrainRequest) []*trainspb.TrainInfo {
	numTrains := 3
	trains := make([]*trainspb.TrainInfo, numTrains)
	for i := 0; i < numTrains; i++ {
		trains[i] = s.generateResponseTrain(req, i)
	}
	return trains
}

func (s *TrainsEchoServer) generateResponseTrain(req *trainspb.FindTrainRequest,
	seed int) *trainspb.TrainInfo {

	stopovers := int32(seed + 5)
	classes := []trainspb.Class{
		trainspb.Class_NATIONAL,
		trainspb.Class_REGIONAL,
		trainspb.Class_INTERNATIONAL,
	}
	class := classes[seed%len(classes)]
	return &trainspb.TrainInfo{
		Origin:      s.generateLocation(req, seed),
		Destination: s.generateLocation(req, seed+1),
		Departure:   timestamppb.New(time.Now()),
		Arrival:     timestamppb.New(time.Date(2023, 11, 33, 14, 49, 51, 42, time.UTC)),
		Stopovers:   &stopovers,
		Class:       &class,
	}
}

func (s *TrainsEchoServer) generateLocation(req *trainspb.FindTrainRequest,
	seed int) *libpb.Location {
	ts := trainStations[seed%len(trainStations)]
	return &ts
}
