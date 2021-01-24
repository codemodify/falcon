package data

type Collection struct {
	ID     string   `json:"id,omitempty"`
	Record []Record `json:"records,omitempty"`
}
