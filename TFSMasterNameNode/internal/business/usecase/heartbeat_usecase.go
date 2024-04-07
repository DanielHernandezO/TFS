package usecase

import (
	"context"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/gateway"
)

type HeartBeatUsecase interface {
	HeartBeat(context *context.Context, socketData *domain.Socket)
}

type heartbeatUsecase struct {
	socketGateway gateway.SocketGateway
}

func NewHeartBeatUsecase(socketGateway gateway.SocketGateway) *heartbeatUsecase {
	return &heartbeatUsecase{
		socketGateway: socketGateway,
	}
}

func (h *heartbeatUsecase) HeartBeat(context *context.Context, socket *domain.Socket) {
	h.socketGateway.Add(socket)
}
