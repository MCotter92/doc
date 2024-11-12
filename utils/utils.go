package utils

type Document struct {
	UUID        string   `json:"UUID"`
	Title       string   `json:"Title"`
	Extension   string   `json:"Extension`
	Location    string   `json:"Location"`
	CreatedDate string   `json:"CreatedDate"`
	Keywords    []string `json:"Keywords"`
}
