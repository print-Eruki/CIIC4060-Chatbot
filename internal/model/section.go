package model

type Section struct {
	Sid      uint64 `json:"sid"`
	Roomid   uint64 `json:"roomid"`
	Mid      uint64 `json:"mid"`
	Cid      uint64 `json:"cid"`
	Semester string `json:"semester"`
	Years    string `json:"years"`
	Capacity uint32 `json:"capacity"`
}
