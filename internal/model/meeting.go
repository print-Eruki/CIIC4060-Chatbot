package model

type Meeting struct {
	Mid       uint64 `json:"mid"`
	Ccode     string `json:"ccode"`
	Starttime string `json:"starttime"` // string for better integration
	Endtime   string `json:"endtime"`
	Cdays     string `json:"cdays"`
}
