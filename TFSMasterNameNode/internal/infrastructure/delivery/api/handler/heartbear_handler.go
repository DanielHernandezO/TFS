package handler

import (
	"context"
	"net/http"

	"github.com/TSF/TFSMasterNameNode/internal/business/constant"
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
)

type heartBeatHandler struct {
	heartBeatUsecase usecase.HeartBeatUsecase
	config.UnimplementedHeartBeatServer
}

func NewHeartBeatHandler(heartBeatUsecase usecase.HeartBeatUsecase) *heartBeatHandler {
	return &heartBeatHandler{
		heartBeatUsecase: heartBeatUsecase,
	}
}

func (h *heartBeatHandler) HeartBeat(context context.Context, socketData *config.Socket) (*config.Response, error) {
	socket := &domain.Socket{
		Ip:   socketData.Ip,
		Port: socketData.Port,
	}

	err := h.heartBeatUsecase.HeartBeat(&context, socket)

	if err != nil {
		return nil, err
	}
	return &config.Response{
		Code:    http.StatusOK,
		Message: constant.AddedServer,
	}, nil
}
