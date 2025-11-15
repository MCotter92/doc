package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

var SetFrontMatterFunc = SetFrontmatter

func TestSetID(t *testing.T) {
	var d Doc
	d.setID()
	if d.Id == uuid.Nil {
		t.Fatal("expected non-nil UUID")
	}
}

func TestSetTitle(t *testing.T) {
	var d Doc
	err := d.setTitle("/tmp/notes/myNote.md")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d.Title != "myNote.md" {
		t.Errorf("expected title=myNote.md, got %s", d.Title)
	}
}

func TestSetDirectory(t *testing.T) {
	tmp := t.TempDir()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"absolute path", filepath.Join(tmp, "myNote.md"), tmp},
		{"relative path", "myNote.md", func() string {
			cwd, _ := os.Getwd()
			return cwd
		}()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Doc
			err := d.setDirectory(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if d.Directory != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, d.Directory)
			}
		})
	}
}

func TestSetPath(t *testing.T) {
	var d Doc
	d.setPath("/tmp/notes", "note.md")
	expected := filepath.Join("/tmp/notes", "note.md")
	if d.Path != expected {
		t.Errorf("expected %q, got %q", expected, d.Path)
	}
}

func TestSetCreatedDate(t *testing.T) {
	var d Doc
	d.setCreatedDate()
	if d.CreatedDate.IsZero() {
		t.Fatal("expected non-zero CreatedDate")
	}
	if time.Since(d.CreatedDate) > time.Second {
		t.Fatal("CreatedDate seems too old")
	}
}

func TestSetKeyword(t *testing.T) {
	var d Doc
	d.setKeyword("golang")
	if d.Keyword != "golang" {
		t.Errorf("expected golang, got %s", d.Keyword)
	}
}

// Failing: rewrite?
func TestNewDoc(t *testing.T) {
	title := "testnote.md"
	keyword := "testing"
	doc, err := NewDoc(title, keyword)

	if err == nil {
		// We expect setUserID to fail because no database exists
		if doc.UserID == uuid.Nil {
			t.Log("expected UserID to be nil due to missing DB — OK")
		} else {
			t.Error("expected UserID to be nil but got a value")
		}
	} else {
		t.Logf("expected error (likely due to DB), got: %v", err)
	}

	if doc.Id == uuid.Nil {
		t.Error("expected non-nil ID")
	}
	if doc.Title != title {
		t.Errorf("expected title=%s, got %s", title, doc.Title)
	}
	if !strings.HasSuffix(doc.Path, title) {
		t.Errorf("expected path to end with %s, got %s", title, doc.Path)
	}
	if doc.Keyword != keyword {
		t.Errorf("expected keyword=%s, got %s", keyword, doc.Keyword)
	}
}

// Failing: expected "frontmatter" got "", stubbing?
func TestCreateDocFile(t *testing.T) {
	tmp := t.TempDir()
	filePath := tmp + "/testDoc.md"

	err := CreateDocFile(filePath, "keyword")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content, _ := os.ReadFile(filePath)
	if string(content) != "frontmatter" {
		t.Errorf("expected 'frontmatter', got %q", string(content))
	}
}
