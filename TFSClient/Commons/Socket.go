package Commons

type Socket struct {
	Ip         string `json:"ip"`
	Port       string `json:"port"`
	Replica_id int32  `json:"replica_id"`
}
