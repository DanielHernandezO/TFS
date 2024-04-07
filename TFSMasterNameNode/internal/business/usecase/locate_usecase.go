package usecase

import (
	"context"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/gateway"
)

type LocateUsecase interface {
	LocateChunk(context *context.Context, location *domain.ChunkLocation) error
}

type locateUsecase struct {
	metadataGateway gateway.MetadataGateway
}

func NewLocateUsecase(metadataGateway gateway.MetadataGateway) *locateUsecase {
	return &locateUsecase{
		metadataGateway: metadataGateway,
	}
}

func (l *locateUsecase) LocateChunk(context *context.Context, location *domain.ChunkLocation) error {
	err := l.metadataGateway.LocateChunk(location)
	if err != nil {
		return err
	}
	return nil
}
