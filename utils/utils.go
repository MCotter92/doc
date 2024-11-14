package utils

import "time"

type Document struct {
	UUID             string    `json:"UUID"`
	Title            string    `json:"Title"`
	Extension        string    `json:"Extension`
	Location         string    `json:"Location"`
	CreatedDate      time.Time `json:"CreatedDate"`
	Keywords         []string  `json:"Keywords"`
	LastModifiedDate time.Time
}
