package model

type User struct {
	Uid        uint64 `json:"uid"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
}
