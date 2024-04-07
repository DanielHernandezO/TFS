package domain

import "time"

type Socket struct {
	Ip   string     `json:"ip"`
	Port string     `json:"port"`
	Date *time.Time `json:"date,omitempty"`
}
