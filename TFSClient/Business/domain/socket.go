package domain

import "time"

type Socket struct {
	Ip        string     `json:"ip"`
	Port      string     `json:"port"`
	ReplicaId int        `json:"replica_id,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
}
