package utils

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
    Id               uuid.UUID `json:"id"`
    Title            string    `json:"title"`
    Extension        string    `json:"extension"`
    Location         string    `json:"location"`
    CreatedDate      time.Time `json:"createdDate"`
    LastModifiedDate time.Time `json:"lastModifiedDate"`
    Keyword          string  `json:"keyword"`
}

type DocumentStore struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Documents []Document `json:"documents"`
}
