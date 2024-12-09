package utils

import "time"

type DocumentStore struct {
	UUID             string    `json:"uuid"`
	CreatedDate      time.Time `json:"CreatedDate"`
	LastModifiedDate time.Time `json:"LastModifiedDate"`
	Keywords         []string  `json:"Keywords"`
	Document         Document  `json:"Document"`
}

type Document struct {
	Title     string `json:"Title"`
	Extension string `json:"Extension"`
	Location  string `json:"Location"`
}
