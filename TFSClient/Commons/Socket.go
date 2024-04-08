package Commons

type Socket struct {
	ip         string `json:"ip"`
	port       string `json:"port"`
	replica_id int    `json:"replica_id"`
}
