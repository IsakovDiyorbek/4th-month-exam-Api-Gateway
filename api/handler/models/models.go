package models

type MediaReq struct {
	Id       string `json:"-"`
	MemoryId string `json:"memory_id"`
	Type     string `json:"type"`
	Url      string `json:"url"`
}
