package data

type Record struct {
	ID     string                 `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}
