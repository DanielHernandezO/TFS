package gateway

import "github.com/TSF/TFSClient/Business/domain"

type NameNodeGateway interface {
	SendMetadata(location string, chunksQuantity int) int
	GetReplicas() ([]domain.Socket, int)
	DeleteMetadata(filename string) int
}
