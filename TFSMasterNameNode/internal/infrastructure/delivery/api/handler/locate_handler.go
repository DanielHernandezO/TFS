package handler

import (
	"context"
	"net/http"

	"github.com/TSF/TFSMasterNameNode/internal/business/constant"
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
)

type locateHandler struct {
	locateUsecase usecase.LocateUsecase
	config.UnimplementedLocateChunkServer
}

func NewLocateHandler(locateUsecase usecase.LocateUsecase) *locateHandler {
	return &locateHandler{
		locateUsecase: locateUsecase,
	}
}

func (l *locateHandler) LocateChunk(context context.Context, chunkLocation *config.ChunkLocation) (*config.Response, error) {
	location := &domain.ChunkLocation{
		Name:      chunkLocation.Name,
		ID:        int(chunkLocation.ChunkId),
		ReplicaId: int(chunkLocation.ReplicaId),
		Socket: domain.Socket{
			Ip:   chunkLocation.Socket.Ip,
			Port: chunkLocation.Socket.Port,
		},
	}

	err := l.locateUsecase.LocateChunk(&context, location)

	if err != nil {
		return &config.Response{}, err
	}
	return &config.Response{
		Code:    http.StatusOK,
		Message: constant.AddedChunk,
	}, nil
}
