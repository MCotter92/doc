package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Extension   string    `json:"extension"`
	Location    string    `json:"location"`
	FullName    string    `json:"FullName"`
	CreatedDate time.Time `json:"createdDate"`
	Keyword     string    `json:"Keyword"`
	Inode       uint64    `json:"inode"`
}

type DocumentStore struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	NotesLocation string
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
	if loc == "." {
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

func (d *Document) SetFullName() {
	d.FullName = d.Location + "/" + d.Title
}

func (d *Document) getFullName() string {
	return d.FullName
}

func (d *Document) SetInode() {
	file := d.FullName
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Printf("Cannot retireve file info: %v\n", err)
	}

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		fmt.Printf("Not a syscall.Stat_t")
	}

	d.Inode = stat.Ino
}

func (d *Document) GetInode() uint64 {
	return d.Inode
}

func (d *Document) SetCreatedDate() {
	d.CreatedDate = time.Now()
}

func (d *Document) GetCreatedDate() time.Time {
	return d.CreatedDate
}

func (d *Document) SetKeyword(keyword string) {
	d.Keyword = keyword
}

func (d *Document) GetKeyword() string {
	return d.Keyword
}
