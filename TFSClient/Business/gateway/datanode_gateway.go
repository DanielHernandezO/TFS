package gateway

import (
	"context"

	"github.com/TSF/TFSClient/Business/domain"
)

type DataNodeGateway interface {
	SendChunk(context context.Context, name string, chunkId int, replicas []domain.Socket, chunk []byte) error
	GetChunk(context context.Context, name string, chunkId int, replicaid int, socket domain.Socket) ([]byte, error)
}
