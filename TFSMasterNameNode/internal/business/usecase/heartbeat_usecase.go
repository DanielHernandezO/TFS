package usecase

import (
	"context"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/gateway"
)

type HeartBeatUsecase interface {
	HeartBeat(context *context.Context, socketData *domain.Socket) error
}

type heartbeatUsecase struct {
	socketGateway gateway.SocketGateway
	fetchGateway  gateway.FetchGateway
}

func NewHeartBeatUsecase(socketGateway gateway.SocketGateway, fetchGateway gateway.FetchGateway) *heartbeatUsecase {
	return &heartbeatUsecase{
		socketGateway: socketGateway,
		fetchGateway:  fetchGateway,
	}
}

func (h *heartbeatUsecase) HeartBeat(context *context.Context, socket *domain.Socket) error {
	h.socketGateway.Add(socket)
	_, err := h.fetchGateway.FetchHeartBeat(context, socket)
	if err != nil {
		return err
	}
	return nil
}
