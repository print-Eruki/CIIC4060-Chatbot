package model

type Room struct {
	Rid         uint64 `json:"rid"`
	Building    string `json:"building"`
	Room_number string `json:"room_number"`
	Capacity    uint32 `json:"capacity"`
}
