package Commons

type NameNodeResponse struct {
	File_name string  `json:"file_name"`
	Chunks    []Chunk `json:"chunks"`
}

type DataNodesList struct {
	Chunks []Chunk `json:"chunks"`
}

type NameNodeRequest struct {
	File_name      string `json:"file_name"`
	ChunksQuantity int    `json:"chunks_quantity"`
}

// Read implements io.Reader.
func (NameNodeRequest) Read(p []byte) (n int, err error) {
	panic("unimplemented")
}
