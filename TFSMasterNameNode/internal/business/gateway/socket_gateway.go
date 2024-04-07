package gateway

import "github.com/TSF/TFSMasterNameNode/internal/business/domain"

type SocketGateway interface {
	Get() map[string]domain.Socket
	Add(socket *domain.Socket)
	Update()
}
