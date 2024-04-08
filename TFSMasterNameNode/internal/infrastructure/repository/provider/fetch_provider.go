package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TSF/TFSMasterNameNode/internal/business/constant"
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var (
	replica = "replica"
)

type fetchProvider struct {
}

func NewFecthProvider() *fetchProvider {
	return &fetchProvider{}
}

func (f *fetchProvider) GetMetadata(context *context.Context) *config.Metadata {
	conn := checkPeerConnection(config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, replica)), config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, replica)))
	if conn == nil {
		panic(constant.FetchError)
	}
	defer conn.Close()
	client := config.NewFetchClient(conn)
	body := &config.VoidRequest{}
	metaData, err := client.FetchLocations(*context, body)
	if err != nil {
		panic(constant.FetchError)
	}
	return metaData
}

func (f *fetchProvider) GetSockets(context *context.Context) *config.Sockets {
	conn := checkPeerConnection(config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, replica)), config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, replica)))
	if conn == nil {
		panic(constant.FetchError)
	}
	defer conn.Close()
	client := config.NewFetchClient(conn)
	body := &config.VoidRequest{}
	sockets, err := client.FetchSocket(*context, body)
	if err != nil {
		panic(constant.FetchError)
	}
	return sockets
}

func (f *fetchProvider) FetchHeartBeat(context *context.Context, socket *domain.Socket) (*config.Response, error) {
	conn := checkPeerConnection(config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, replica)), config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, replica)))
	if conn == nil {
		return nil, nil
	}
	defer conn.Close()
	client := config.NewFetchClient(conn)
	body := &config.Socket{
		Ip:   socket.Ip,
		Port: socket.Port,
	}
	message, err := client.FetchHeartBeat(*context, body)
	if err != nil {
		return nil, fmt.Errorf(constant.CouldntUpdateReplica)
	}
	return message, nil
}

func (f *fetchProvider) FetchLocation(context *context.Context, location *domain.ChunkLocation) (*config.Response, error) {
	conn := checkPeerConnection(config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, replica)), config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, replica)))
	if conn == nil {
		return nil, nil
	}
	defer conn.Close()
	client := config.NewFetchClient(conn)
	body := &config.ChunkLocation{
		Name:      location.Name,
		ChunkId:   int32(location.ID),
		ReplicaId: int32(location.ReplicaId),
		Socket: &config.Socket{
			Ip:        location.Socket.Ip,
			Port:      location.Socket.Port,
			ReplicaId: int32(location.Socket.ReplicaId),
		},
	}
	message, err := client.FetchLocateChunk(*context, body)
	if err != nil {
		return nil, fmt.Errorf(constant.CouldntUpdateReplica)
	}
	return message, nil
}

func (f *fetchProvider) FetchMetadata(context *context.Context, metadata *domain.Metadata) (*config.Response, error) {
	conn := checkPeerConnection(config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, replica)), config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, replica)))
	if conn == nil {
		return nil, nil
	}
	defer conn.Close()
	client := config.NewFetchClient(conn)
	body := &config.FileMetadataFetch{
		FileId:         metadata.Name,
		ChunksQuantity: int32(metadata.ChunksQuantity),
	}
	message, err := client.FetchMetadata(*context, body)
	if err != nil {
		return nil, fmt.Errorf(constant.CouldntUpdateReplica)
	}
	return message, nil
}

func (f *fetchProvider) DeleteMetadataFetch(context *context.Context, fileName string) (*config.Response, error) {
	conn := checkPeerConnection(config.GetStringPropetyBykey(fmt.Sprintf(config.IpConfig, replica)), config.GetStringPropetyBykey(fmt.Sprintf(config.PortConfig, replica)))
	if conn == nil {
		return nil, nil
	}
	defer conn.Close()
	client := config.NewFetchClient(conn)
	body := &config.DeleteMetadata{
		FileId: fileName,
	}
	message, err := client.DeleteMetadataFetch(*context, body)
	if err != nil {
		return nil, fmt.Errorf(constant.CouldntUpdateReplica)
	}
	return message, nil
}

func checkPeerConnection(ip string, port string) *grpc.ClientConn {

	address := fmt.Sprintf("%s:%s", ip, port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("error conecting to:%s", address)
		return nil
	}

	timeout := 15 * time.Second
	deadline := time.Now().Add(timeout)
	for {
		state := conn.GetState()
		if state == connectivity.Ready {
			return conn
		}

		if time.Now().After(deadline) {
			return nil
		}

		time.Sleep(time.Second)
	}
}
