package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"services/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {

}



func (*server) Greet (ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("Function called... ")

	fistName:= req.GetGreeting().GetFistName()
	lastName:=req.GetGreeting().GetLastName()

	result := "Hello : " +fistName +", " + lastName

	res := &greetpb.GreetResponse {
			Result: result,
		}
	return res, nil

}

func main() {
	fmt.Println("Hello World From Server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v: ", err)
	}

	s:= grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if s.Serve(lis); err != nil  {
		log.Fatalf("failed to Serve %v", err)
	}

}