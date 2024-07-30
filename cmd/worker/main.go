package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pbw "github.com/chkda/mapreduce/rpc/worker"
	"github.com/chkda/mapreduce/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	workerId := flag.String("workerid", "", "WorkerId")
	ip := flag.String("ip", "localhost", "Worker IP Addr")
	masterIP := flag.String("master", "localhost:4000", "Master IP Addr")
	port := flag.String("port", "5000", "Worker GRPC port")
	flag.Parse()
	if *masterIP == "" || *workerId == "" || *port == "" {
		panic("null values")
	}
	config := &worker.Config{
		GRPCPort: *port,
		MasterIP: *masterIP,
		IP:       *ip + ":" + *port,
	}
	mrWorker, err := worker.New(config)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	pbw.RegisterWorkerServer(server, mrWorker)
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
