package api

import (
	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/delivery/api/handler"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/repository/metadata"
	socket "github.com/TSF/TFSMasterNameNode/internal/infrastructure/repository/server"
)

type Dependencies struct {
	pingHandler        handler.PingHandller
	metadataHandler    handler.MetadataHandler
	locateChunkHandler config.LocateChunkServer
	heartBeatHandler   config.HeartBeatServer
	socketHandler      handler.SocketHandler
	fetchHandler       config.FetchServer
}

func buildDependencies() *Dependencies {
	//repository
	metadataGateway := metadata.NewMetadataRepository()
	socketGateway := socket.NewSocketRepository()

	//usecase
	metadataUsecase := usecase.NewMetadataUsecase(metadataGateway)
	locateUsecase := usecase.NewLocateUsecase(metadataGateway)
	heartBeatusecase := usecase.NewHeartBeatUsecase(socketGateway)
	socketUsecase := usecase.NewSocketUsecase(socketGateway)
	fetchUsecase := usecase.NewFetchUsecase(socketGateway, metadataGateway)

	return &Dependencies{
		pingHandler:        handler.NewPingHanlder(),
		metadataHandler:    handler.NewMetadataHanlder(metadataUsecase),
		locateChunkHandler: handler.NewLocateHandler(locateUsecase),
		heartBeatHandler:   handler.NewHeartBeatHandler(heartBeatusecase),
		socketHandler:      handler.NewSocketHandler(socketUsecase),
		fetchHandler:       handler.NewFetchHandler(fetchUsecase),
	}
}
