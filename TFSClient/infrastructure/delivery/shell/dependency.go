package shell

import (
	"github.com/TSF/TFSClient/Business/usecase"
	"github.com/TSF/TFSClient/infrastructure/delivery/shell/handler"
	"github.com/TSF/TFSClient/infrastructure/repository/provider"
)

type Dependencies struct {
	writeHandler handler.WriteHandler
	readHandler  handler.ReadHandler
}

func buildDependencies() *Dependencies {

	//provider
	nameNodeGateway := provider.NewNameNodeProvider()
	dataNodeGateway := provider.NewDataNodeProvider()

	//usecase
	writeUsecase := usecase.NewWriteUsecase(nameNodeGateway, dataNodeGateway)
	readUsecase := usecase.NewReadUsecase(nameNodeGateway, dataNodeGateway)

	return &Dependencies{
		writeHandler: handler.NewWriteHandler(writeUsecase),
		readHandler:  handler.NewReadHandler(readUsecase),
	}
}
