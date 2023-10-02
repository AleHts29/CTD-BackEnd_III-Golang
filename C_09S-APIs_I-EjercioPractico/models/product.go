package models

import "time"

type Product struct {
	Id         int       `json:"Id"`
	Name       string    `json:"Name"`
	Price      float64   `json:"Price"`
	Stock      int       `json:"Stock"`
	Code       string    `json:"Code"`
	Publish    bool      `json:"Publish"`
	CreateDate time.Time `json:"CreateDate"`
}
