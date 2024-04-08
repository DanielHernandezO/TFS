package Commons

type DataNodeResponse struct {
	Block []byte `json:"block"`
}

type DataNodeRequest struct {
	BlockID  int    `json:"number"`
	Filename string `json:"string"`
}
