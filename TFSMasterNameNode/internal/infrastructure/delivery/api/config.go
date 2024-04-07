package api

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/delivery"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	grpcService = "grpc"
	restService = "rest"
)

type api struct{}

func NewApi() delivery.Strategy {
	return &api{}
}

func (r *api) Start() {
	config.LoadConfigs()
	dependencies := buildDependencies()

	lis, err := net.Listen(config.GetStringPropetyBykey(config.Network), fmt.Sprintf(":%s", config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, grpcService))))
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegister := grpc.NewServer()
	config.RegisterLocateChunkServer(serverRegister, dependencies.locateChunkHandler)
	config.RegisterHeartBeatServer(serverRegister, dependencies.heartBeatHandler)
	config.RegisterFetchServer(serverRegister, dependencies.fetchHandler)
	log.Printf("gRPC server listening on :%s", config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, grpcService)))
	go serverRegister.Serve(lis)

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		go func() {
			for range ticker.C {
				dependencies.socketHandler.Update()
			}
		}()
	}()

	router := gin.Default()
	mapUrls(router, dependencies)
	router.Run(fmt.Sprintf(":%s", config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, restService))))
}
