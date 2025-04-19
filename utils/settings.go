package utils

import "github.com/google/uuid"

type Profile struct {
	Id            uuid.UUID
	Editor        string
	NotesLocation string
}

func (p *Profile) NewProfile(editor, NotesLocation string) {
	p.Id = uuid.New()
	p.Editor = editor
	p.NotesLocation = NotesLocation
}

func (p *Profile) SetEditor(editor string) {
	p.Editor = editor
}

func (p *Profile) SetNotesLocation(NotesLocation string) {
	p.NotesLocation = NotesLocation
}
