package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/aofiee/grpcservices/protos/person"
	"google.golang.org/grpc"
)

type personData struct {
	person.UnimplementedPersonServer
}
type employee struct {
	ID       string
	FullName string
	Age      int32
	Email    string
	Address  string
}

type myEmployee struct {
	persons []employee
}

var e = myEmployee{}

func (p *personData) Person(ctx context.Context, in *person.GetPersonFromIDRequest) (*person.PersonResponse, error) {
	log.Println("Receive : ", in.ID)
	for _, v := range e.persons {
		if v.ID == in.ID {
			return &person.PersonResponse{
				ID:       v.ID,
				FullName: v.FullName,
				Age:      v.Age,
				Email:    v.Email,
				Address:  v.Address,
			}, nil
		}
	}
	return nil, errors.New("can't find id")
}

func (p *personData) AllPerson(ctx context.Context, in *person.GetALLPersonsRequest) (*person.AllPersonResponse, error) {
	emp := []*person.PersonResponse{}
	for _, v := range e.persons {
		emp = append(emp, &person.PersonResponse{
			ID:       v.ID,
			FullName: v.FullName,
			Age:      v.Age,
			Email:    v.Email,
			Address:  v.Address,
		})
	}
	return &person.AllPersonResponse{Persons: emp}, nil
}
func main() {
	log.SetFlags(log.Ltime)
	e.persons = append(e.persons, employee{
		ID:       "2779dd47-82d3-4548-bfbe-1a5374954436",
		FullName: "Khomkrid Lerdprasert",
		Age:      40,
		Email:    "aofieee666@gmail.com",
		Address:  "44/261",
	})
	e.persons = append(e.persons, employee{
		ID:       "4b59f7f2-3b38-4098-834c-1763815fe273", //uuid.New().String(),
		FullName: "Arnon Kidlerdpol",
		Age:      35,
		Email:    "snappytux@gmail.com",
		Address:  "57/1",
	})
	e.persons = append(e.persons, employee{
		ID:       "be0d794b-3893-4159-8aac-96bd61045468",
		FullName: "Wittaya Jitsuwan",
		Age:      32,
		Email:    "maxzerocity@gmail.com",
		Address:  "331",
	})
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	per := &personData{}
	person.RegisterPersonServer(s, per)
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to person serve: %v", err)
	}
}
