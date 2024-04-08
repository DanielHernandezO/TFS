package provider

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TSF/TFSClient/Business/domain"
	"github.com/TSF/TFSClient/infrastructure/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type dataNodeProvider struct {
}

func NewDataNodeProvider() *dataNodeProvider {
	return &dataNodeProvider{}
}

func (d *dataNodeProvider) SendChunk(context context.Context, name string, chunkId int, replicas []domain.Socket, chunk []byte) error {
	conn := checkPeerConnection(replicas[0].Ip, replicas[0].Port)
	if conn == nil {
		return fmt.Errorf("error connecting dataNode")
	}
	defer conn.Close()
	client := config.NewDataNodeServiceClient(conn)
	replicas = replicas[1:]

	dataNodeSockets := []*config.DataNodeSocket{}
	for _, replica := range replicas {
		dataNodeSockets = append(dataNodeSockets, &config.DataNodeSocket{
			Ip:   replica.Ip,
			Port: replica.Port,
		})
	}
	body := &config.ChunkWriteRequest{
		ChunkMetadata: &config.ChunkMetadata{
			Name:      name,
			ChunkId:   int32(chunkId),
			ReplicaId: 0,
		},
		PipelineDataNodes: dataNodeSockets,
		ChunkBytes:        chunk,
	}
	_, err := client.WriteChunk(context, body)
	if err != nil {
		return err
	}
	return nil
}

func checkPeerConnection(ip string, port string) *grpc.ClientConn {

	address := fmt.Sprintf("%s:%s", ip, port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("error conecting to:%s", address)
		return nil
	}

	timeout := 3 * time.Second
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
