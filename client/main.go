package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aofiee/grpcservices/protos/person"
	"github.com/aofiee/grpcservices/protos/position"
	"github.com/fatih/color"

	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Ltime)
	if arg() {
		switch os.Args[1] {
		case "-a":
			clientConn, err := grpc.DialContext(context.Background(), ":50051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("failed to dail: %v", err)
			}
			client := person.NewPersonClient(clientConn)
			res, err := client.AllPerson(context.Background(), &person.GetALLPersonsRequest{})
			if err != nil {
				log.Fatalf("failed to say: %v", err)
			}
			for _, v := range res.Persons {
				log.Println(" : ", v.ID)
				log.Println(" : ", v.FullName)
				log.Println(" : ", v.Age)
				log.Println(" : ", v.Email)
				log.Println(" : ", v.Address)
				log.Println("-----------------------")
			}
			break
		case "-uuid":
			clientConn, err := grpc.DialContext(context.Background(), ":50051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("failed to dail: %v", err)
			}
			client := person.NewPersonClient(clientConn)
			res, err := client.Person(context.Background(), &person.GetPersonFromIDRequest{ID: os.Args[2]})
			if err != nil {
				log.Fatalf("failed to say: %v", err)
			}
			log.Println(" : ", res.ID)
			log.Println(" : ", res.FullName)
			log.Println(" : ", res.Age)
			log.Println(" : ", res.Email)
			log.Println(" : ", res.Address)
			break
		case "-puuid":
			clientConn, err := grpc.DialContext(context.Background(), ":50052", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("failed to dail: %v", err)
			}
			client := position.NewPositionClient(clientConn)
			res, err := client.Position(context.Background(), &position.UserPositionFromUUIDRequest{UUID: os.Args[2]})
			if err != nil {
				log.Fatalf("failed to say: %v", err)
			}
			log.Println(" : ", res.FullName)
			log.Println(" : ", res.Email)
			log.Println(" : ", res.Position)
			log.Println(" : ", res.JobDesc)
			break
		case "-ap":
			clientConn, err := grpc.DialContext(context.Background(), ":50052", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("failed to dail: %v", err)
			}
			client := position.NewPositionClient(clientConn)
			res, err := client.AllPosition(context.Background(), &position.AllPositionsRequest{})
			if err != nil {
				log.Fatalf("failed to say: %v", err)
			}
			for _, v := range res.Positions {
				log.Println(" : ", v.FullName)
				log.Println(" : ", v.Email)
				log.Println(" : ", v.Position)
				log.Println(" : ", v.JobDesc)
				log.Println("------------------------------------------------------------------------------------------")
			}

			break
		}

	}

}
func arg() bool {
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	if len(os.Args) < 2 {
		fmt.Println(blue("-a for all records."))
		fmt.Println(blue("-uuid string for one person."))
		fmt.Println(blue("-ap list all position."))
		fmt.Println(blue("-puuid position from user uuid."))
		os.Exit(1)
	}
	return true
}
