package socket

import (
	"sync"
	"time"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
)

var (
	minutesHeartBeatLimit = 5
)

type socketRepository struct {
	sockets map[string]domain.Socket
	mu      sync.Mutex
}

func NewSocketRepository() *socketRepository {
	return &socketRepository{
		sockets: map[string]domain.Socket{},
	}
}

func (s *socketRepository) Get() map[string]domain.Socket {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.sockets
}

func (s *socketRepository) Add(socket *domain.Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	socket.Date = &now
	s.sockets[socket.Ip] = *socket
}

func (s *socketRepository) Update() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key, value := range s.sockets {
		difference := time.Since(*value.Date)
		minutes := int(difference.Minutes())
		if minutes > minutesHeartBeatLimit {
			delete(s.sockets, key)
		}
	}
}
