package api

import (
	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/delivery/api/handler"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/repository/metadata"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/repository/provider"
	socket "github.com/TSF/TFSMasterNameNode/internal/infrastructure/repository/server"
)

type Dependencies struct {
	pingHandler         handler.PingHandller
	metadataHandler     handler.MetadataHandler
	locateChunkHandler  config.LocateChunkServer
	heartBeatHandler    config.HeartBeatServer
	socketHandler       handler.SocketHandler
	fetchHandler        config.FetchServer
	executeFetchHandler handler.ExecuteFetchHandler
}

func buildDependencies() *Dependencies {
	//repository
	metadataGateway := metadata.NewMetadataRepository()
	socketGateway := socket.NewSocketRepository()
	fetchGateway := provider.NewFecthProvider()

	//usecase
	metadataUsecase := usecase.NewMetadataUsecase(metadataGateway, fetchGateway)
	locateUsecase := usecase.NewLocateUsecase(metadataGateway, fetchGateway)
	heartBeatusecase := usecase.NewHeartBeatUsecase(socketGateway, fetchGateway)
	socketUsecase := usecase.NewSocketUsecase(socketGateway)
	fetchUsecase := usecase.NewFetchUsecase(socketGateway, metadataGateway, fetchGateway)

	return &Dependencies{
		pingHandler:         handler.NewPingHanlder(),
		metadataHandler:     handler.NewMetadataHanlder(metadataUsecase),
		locateChunkHandler:  handler.NewLocateHandler(locateUsecase),
		heartBeatHandler:    handler.NewHeartBeatHandler(heartBeatusecase),
		socketHandler:       handler.NewSocketHandler(socketUsecase),
		fetchHandler:        handler.NewFetchHandler(fetchUsecase),
		executeFetchHandler: handler.NewExecuteFetchHandler(fetchUsecase),
	}
}
