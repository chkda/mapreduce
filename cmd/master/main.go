package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/chkda/mapreduce/master"
	pbm "github.com/chkda/mapreduce/rpc/master"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var FILE_LOC = "/config/master/config.json"

func main() {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := currDir + FILE_LOC
	file, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	config := &master.Config{}
	err = json.Unmarshal(fileBytes, config)
	if err != nil {
		panic(err)
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
