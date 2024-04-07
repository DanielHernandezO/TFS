package usecase

import (
	"math/rand"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/gateway"
)

var (
	replicas = 3
)

type SokcetUsecase interface {
	GetReplicas() *[]domain.Socket
	Update()
}

type socketUsecase struct {
	socketGateway gateway.SocketGateway
}

func NewSocketUsecase(socketGateway gateway.SocketGateway) *socketUsecase {
	return &socketUsecase{
		socketGateway: socketGateway,
	}
}

func (s *socketUsecase) GetReplicas() *[]domain.Socket {
	sockets := s.socketGateway.Get()

	socketList := make([]domain.Socket, 0, len(sockets))
	for key, value := range sockets {
		socketList = append(socketList, domain.Socket{Ip: key, Port: value.Port})
	}

	rand.Shuffle(len(socketList), func(i, j int) {
		socketList[i], socketList[j] = socketList[j], socketList[i]
	})

	selectedSockets := []domain.Socket{}
	for i := 0; i < replicas && i < len(socketList); i++ {
		index := rand.Intn(len(socketList) - i)
		selectedSockets = append(selectedSockets, socketList[index])

		socketList = append(socketList[:index], socketList[index+1:]...)
	}

	return &selectedSockets
}

func (s *socketUsecase) Update() {
	s.socketGateway.Update()
}
