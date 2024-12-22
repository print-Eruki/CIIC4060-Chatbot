package model

import "time"

type User struct {
	Uid        uint64    `json:"uid"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
}
