package main

import (
	"log"
	"net"

	"github.com/phayes/freeport"
	"google.golang.org/grpc"

	pb "github.com/getstackhead/pluginlib/gen"
)

var Plugin MyPluginWrapper

type MyPluginWrapper struct {
}

type plugin struct {
	pb.UnimplementedStackHeadPlugin
}

func (p MyPluginWrapper) LaunchServer(callback func(int)) {
	// Set up connection with local StackHead server
	conn, err := grpc.Dial("localhost:1412", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	StackHeadClient := pb.NewStackHeadServiceClient(conn)

	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
		return
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStackHeadPluginServer(s, &plugin{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	callback(port)
}
