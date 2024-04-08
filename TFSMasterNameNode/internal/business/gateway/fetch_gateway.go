package gateway

import (
	"context"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
)

type FetchGateway interface {
	GetMetadata(context *context.Context) *config.Metadata
	GetSockets(context *context.Context) *config.Sockets
	FetchHeartBeat(context *context.Context, socket *domain.Socket) (*config.Response, error)
	FetchLocation(context *context.Context, location *domain.ChunkLocation) (*config.Response, error)
	FetchMetadata(context *context.Context, metadata *domain.Metadata) (*config.Response, error)
	DeleteMetadataFetch(context *context.Context, fileName string) (*config.Response, error)
}
