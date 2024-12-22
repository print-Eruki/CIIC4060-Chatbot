package model

type Requisite struct {
	Classid uint64 `json:"classid"`
	Reqid   uint64 `json:"reqid"`
	Prereq  bool   `json:"prereq"`
}
