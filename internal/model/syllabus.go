package model

type Syllabus struct {
	Chunkid        uint64 `json:"chunkid"`
	Courseid       uint64 `json:"courseid"`
	Embedding_text string `json:"embedding_text"`
	Chunk          string `json:"chunk"` // Might need to convert to ts_vector in the sql query
}
