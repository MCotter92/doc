package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type User struct {
	UserName      string `yaml:"user_name"`
	NotesLocation string `yaml:"notes_location"`
	Editor        string `yaml:"editor"`
	ConfigPath    string `yaml:"config_path"`
}

// NewUser creates a new user by prompting for all required information
func NewUser() (*User, error) {
	user := &User{}

	// Set config path first (doesn't require user input)
	if err := user.setDefaultConfigPath(); err != nil {
		return nil, fmt.Errorf("failed to set config path: %w", err)
	}

	if err := user.setDefaultNotesLocation(); err != nil {
		return nil, fmt.Errorf("failed to set notes location: %w", err)
	}

	// Prompt for user information
	if err := user.promptUserName(); err != nil {
		return nil, err
	}

	if err := user.promptEditor(); err != nil {
		return nil, err
	}

	return user, nil
}

// Private methods for setting up the user
func (u *User) setDefaultConfigPath() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	u.ConfigPath = filepath.Join(homeDir, ".config", "doc", "userConfig.yaml")
	return nil
}

func (u *User) setDefaultNotesLocation() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	u.NotesLocation = path
	return nil
}

func (u *User) promptUserName() error {
	fmt.Print("What is your user name? ")
	input, err := readInput()
	if err != nil {
		return err
	}
	u.UserName = input
	return nil
}

func (u *User) promptEditor() error {
	fmt.Print("What editor do you want to use by default? (e.g., nvim, code, emacs): ")
	input, err := readInput()
	if err != nil {
		return err
	}
	u.Editor = input
	return nil
}

// Helper function to read and clean user input
func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
