package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Document struct {
    Id               uuid.UUID `json:"id"`
    Title            string    `json:"title"`
    Extension        string    `json:"extension"`
    Location         string    `json:"location"`
    CreatedDate      time.Time `json:"createdDate"`
    Keyword          string    `json:"Keyword"`
}

type DocumentStore struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Documents []Document `json:"documents"`
}

type iDoc interface {
    setID()
    getID()
    setTitle()
    getTitle()
    setExtension()
    getExtension()
    setLocation()
    getLocation()
    setCreatedDate()
    getCreatedDate()
    setLastModifiedDate()
    getLastModifiedDate()
    setKeyword()
    getKeyword()
}

type iStore interface {
    unmarshalJson()
    marshalJson()
}

func (d *Document) SetID() {
    d.Id = uuid.New()
}

func (d *Document) GetID() uuid.UUID {
    return d.Id
}

func (d *Document) SetTitle(fileName string) {
    stats, err := os.Stat(fileName)
    if err != nil {
        fmt.Println("Cannot get os.Stats(fileName)", err)
    }
    d.Title = stats.Name()
}

func (d *Document) GetTitle() string {
    return d.Title
}

func (d *Document) SetExtension(fileName string) {
    ext := filepath.Ext(fileName)
    d.Extension = ext
}

func (d *Document) GetExtension() string {
    return d.Extension
}

func (d *Document) SetLocation(fileName string) {
    loc := filepath.Dir(fileName)
    if loc == "."{
        _loc, err := os.Getwd()
        if err != nil {
            fmt.Printf("Cannot retrieve file location: %v\n", err)
        }
        loc = _loc
    }
    d.Location = loc
}

func (d *Document) GetLocation() string {
    return d.Location
}

func (d *Document) SetCreatedDate() {
    d.CreatedDate = time.Now()
}

func (d *Document) GetCreatedDate() time.Time{
    return d.CreatedDate 
}

func (d *Document) SetKeyword(keyword string) {
   d.Keyword =  keyword 
}

func (d *Document) GetKeyword() string {
    return d.Keyword
}
