package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/aofiee/grpcservices/protos/person"
	"github.com/aofiee/grpcservices/protos/position"
	"google.golang.org/grpc"
)

type positionData struct {
	position.UnimplementedPositionServer
}

type positionOfEmployee struct {
	UUID     string
	Position string
	JobDesc  string
}

type organizationPosition struct {
	positions []positionOfEmployee
}

var o = organizationPosition{}

func (p *positionData) Position(ctx context.Context, in *position.UserPositionFromUUIDRequest) (*position.PositionResponse, error) {
	log.Println("Receive : ", in.UUID)
	for _, v := range o.positions {
		if v.UUID == in.UUID {
			clientConn, err := grpc.DialContext(context.Background(), ":50051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("failed to dail: %v", err)
			}
			client := person.NewPersonClient(clientConn)
			res, err := client.Person(context.Background(), &person.GetPersonFromIDRequest{ID: in.UUID})
			return &position.PositionResponse{
				FullName: res.FullName,
				Email:    res.Email,
				Position: v.Position,
				JobDesc:  v.JobDesc,
			}, nil
		}
	}
	return nil, errors.New("can't find id")
}
func (p *positionData) AllPosition(ctx context.Context, in *position.AllPositionsRequest) (*position.AllPositionsResponse, error) {
	pos := []*position.PositionResponse{}
	for _, v := range o.positions {
		clientConn, err := grpc.DialContext(context.Background(), ":50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("failed to dail: %v", err)
		}
		client := person.NewPersonClient(clientConn)
		res, err := client.Person(context.Background(), &person.GetPersonFromIDRequest{ID: v.UUID})
		pos = append(pos, &position.PositionResponse{
			FullName: res.FullName,
			Email:    res.Email,
			Position: v.Position,
			JobDesc:  v.JobDesc,
		})
	}
	return &position.AllPositionsResponse{Positions: pos}, nil
}
func main() {
	log.SetFlags(log.Ltime)
	o.positions = append(o.positions, positionOfEmployee{
		UUID:     "2779dd47-82d3-4548-bfbe-1a5374954436",
		Position: "Mobile Developer",
		JobDesc:  "A quick look at the mobile app developer job description and duties reveals very slight differences, if any.",
	})
	o.positions = append(o.positions, positionOfEmployee{
		UUID:     "4b59f7f2-3b38-4098-834c-1763815fe273",
		Position: "Operation Developer",
		JobDesc: `THE ROLE WILL CONSIST OF:
		Track leading indicators of renewals and upsell and reporting on results.
		Monitor the timing and content of touch points of the Customer Care Team to drive optimal adoption and CSAT score.
		Define processes to detect early signals of at-risk renewals.
		Define processes to help identify top customers for upsell.
		Together with Revenue Enablement, develop curriculum material that can be useful to the Customer Care Team.
		Implementing and managing software that facilitates CSM, RMs and CSEs operational activities and efficiencies.`,
	})
	o.positions = append(o.positions, positionOfEmployee{
		UUID:     "4b59f7f2-3b38-4098-834c-1763815fe273",
		Position: "Operation Developer",
		JobDesc: `THE ROLE WILL CONSIST OF:
		Track leading indicators of renewals and upsell and reporting on results.
		Monitor the timing and content of touch points of the Customer Care Team to drive optimal adoption and CSAT score.
		Define processes to detect early signals of at-risk renewals.
		Define processes to help identify top customers for upsell.
		Together with Revenue Enablement, develop curriculum material that can be useful to the Customer Care Team.
		Implementing and managing software that facilitates CSM, RMs and CSEs operational activities and efficiencies.`,
	})
	o.positions = append(o.positions, positionOfEmployee{
		UUID:     "be0d794b-3893-4159-8aac-96bd61045468",
		Position: "Data Science",
		JobDesc: `Typically a data scientist needs to:

		Use strong business acumen, as well as an ability to communicate findings, and mine vast amounts of data for useful insights
		Use these insights to influence how an organisation approaches business challenges
		Use a combined knowledge of computer science and applications, modelling, statistics, analytics and maths to solve problems
		Extract data from multiple sources
		Sift and analyse data from multiple angles, looking for trends that highlight problems or opportunities
		Communicate important information and insights to business and IT leaders
		Make recommendations to adapt existing business strategies`,
	})

	l, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pos := &positionData{}
	position.RegisterPositionServer(s, pos)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to position serve: %v", err)
	}
}
