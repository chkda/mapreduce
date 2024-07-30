package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/chkda/mapreduce/master"
	pbm "github.com/chkda/mapreduce/rpc/master"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.String("port", "4000", "Master GRPC port")
	filename := flag.String("filename", "", "File to perform task")
	numReduce := flag.Int("numreduce", 0, "No of reduce workers")
	deadWorkersThreshold := flag.Int("dead-workers-threshold", 1, "Dead workers threshold")
	flag.Parse()
	if *filename == "" || *numReduce == 0 {
		panic("null values")
	}
	config := &master.Config{
		GRPCPort:             *port,
		Filename:             *filename,
		NumReduce:            *numReduce,
		DeadWorkersThreshold: *deadWorkersThreshold,
	}
	mrMaster := master.New(config)

	server := grpc.NewServer()
	pbm.RegisterMasterServer(server, mrMaster)
	reflection.Register(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort))
	if err != nil {
		panic(err)
	}
	log.Println("Starting GRPC Server...")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}

}
