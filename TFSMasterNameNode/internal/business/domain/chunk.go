package domain

type ChunkLocation struct {
	Name      string
	ID        int
	ReplicaId int
	Socket    Socket
}
