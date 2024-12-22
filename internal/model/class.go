package model

type Class struct {
	Cid       uint64 `json:"cid"`
	Ccode     string `json:"ccode"`
	Cname     string `json:"cname"`
	Cred      int32  `json:"cred"`
	Cdesc     string `json:"cdesc"`
	Csyllabus string `json:"csyllabus"`
	Term      string `json:"term"`
	Years     string `json:"years"`
}
