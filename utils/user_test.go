package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func tempHome(t *testing.T) string {
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)

	return tmpDir
}

func TestSetUserID(t *testing.T) {
	var u User
	u.setUserID()

	if u.ID == uuid.Nil {
		t.Fatal("Expected uuid to be set, got nil UUID.")
	}
}

func TestSetDefaultConfigPath(t *testing.T) {
	tmpHome := t.TempDir()

	oldGetUserHomeDir := getUserHomeDir
	getUserHomeDir = func() (string, error) { return tmpHome, nil }
	defer func() { getUserHomeDir = oldGetUserHomeDir }()

	var u User
	err := u.setDefaultConfigPath()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := filepath.Join(tmpHome, "Documents", "Notes")
	if u.ConfigPath != expected {
		t.Errorf("expected config path %q, got %q", expected, u.ConfigPath)
	}
}

func TestSetDefaultNotesLocation(t *testing.T) {
	var u User
	err := u.setDefaultNotesLocation()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cwd, _ := os.Getwd()
	if u.NotesLocation != cwd {
		t.Errorf("expected notesLocation=%q, got %q", cwd, u.NotesLocation)
	}
}

func TestValidateCreatesNotesDirectory(t *testing.T) {
	tmp := t.TempDir()
	user := &User{
		UserName:      "bob",
		Editor:        "nvim",
		NotesLocation: filepath.Join(tmp, "nonexistent"),
	}

	err := user.Validate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := os.Stat(user.NotesLocation); os.IsNotExist(err) {
		t.Fatalf("expected directory to be created")
	}
}

func TestValidateErrorsOnEmptyFields(t *testing.T) {
	tests := []struct {
		name string
		user User
		want string
	}{
		{"missing username", User{}, "userName cannot be empty"},
		{"missing editor", User{UserName: "bob"}, "editor cannot be empty"},
		{"missing notes", User{UserName: "bob", Editor: "vim"}, "notesLocation cannot be empty"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.user.Validate()
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("expected error %q, got %v", test.want, err)
			}
		})
	}
}

func TestPromptUserName(t *testing.T) {
	r, w, _ := os.Pipe()
	fmt.Fprint(w, "alice\n")
	w.Close()

	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = r

	var u User
	err := u.promptUserName()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.UserName != "alice" {
		t.Errorf("expected username alice, got %q", u.UserName)
	}
}

func TestPromptEditor_DefaultsToNvim(t *testing.T) {
	r, w, _ := os.Pipe()
	fmt.Fprint(w, "\n") // simulate pressing Enter
	w.Close()

	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = r

	var u User
	err := u.promptEditor()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Editor != "nvim" {
		t.Errorf("expected editor=nvim, got %q", u.Editor)
	}
}

func TestUpdateConfigFile(t *testing.T) {
	viper.Reset()
	tmpHome := tempHome(t)
	user := &User{
		ID:            uuid.New(),
		UserName:      "bob",
		Editor:        "nvim",
		NotesLocation: tmpHome,
		ConfigPath:    filepath.Join(tmpHome, "config.yaml"),
	}

	err := user.UpdateConfigFile("userName", "alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.UserName != "alice" {
		t.Errorf("expected userName=alice, got %q", user.UserName)
	}
}

func TestSaveConfigFileCreatesFile(t *testing.T) {
	viper.Reset()
	tmp := t.TempDir()
	configPath := filepath.Join(tmp, "conf", "user.yaml")

	user := &User{
		ID:            uuid.New(),
		UserName:      "bob",
		Editor:        "vim",
		NotesLocation: tmp,
		ConfigPath:    configPath,
	}

	if err := user.saveConfigFile(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatalf("expected config file to exist at %s", configPath)
	}
}

func TestReadInput(t *testing.T) {
	r, w, _ := os.Pipe()
	fmt.Fprint(w, "hello world\n")
	w.Close()

	old := os.Stdin
	defer func() { os.Stdin = old }()
	os.Stdin = r

	got, err := readInput()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hello world" {
		t.Errorf("expected 'hello world', got %q", got)
	}
}
