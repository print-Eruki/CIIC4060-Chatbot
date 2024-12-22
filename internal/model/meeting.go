package model

import "time"

type Meeting struct {
	Mid       uint64    `json:"mid"`
	Ccode     string    `json:"ccode"`
	Starttime time.Time `json:"starttime"`
	Endtime   time.Time `json:"endtime"`
	Cdays     string    `json:"cdays"`
}
