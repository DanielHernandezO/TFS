package Commons

type Chunk struct {
	ChunkId int    `json:"chunk_id"`
	Sockets Socket `json:"sockets"`
}
