package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

type User struct {
	ID            string `yaml:"ID"`
	UserName      string `yaml:"userName"`
	NotesLocation string `yaml:"notesLocation"`
	Editor        string `yaml:"editor"`
	ConfigPath    string `yaml:"configPath"`
}

func NewUser() (*User, error) {
	user := &User{}

	user.setUserID()

	if err := user.setDefaultConfigPath(); err != nil {
		return nil, fmt.Errorf("failed to set config path: %w", err)
	}

	if err := user.setDefaultNotesLocation(); err != nil {
		return nil, fmt.Errorf("failed to set notes location: %w", err)
	}

	// Prompt for user information
	if err := user.promptUserName(); err != nil {
		return nil, fmt.Errorf("failed to set user name: %w", err)
	}

	if err := user.promptEditor(); err != nil {
		return nil, fmt.Errorf("failed to set editor: %w", err)
	}

	return user, nil
}

func (u *User) setUserID() {
	u.ID = uuid.New().String()
}

func (u *User) setDefaultConfigPath() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to set config path: %w", err)
	}
	u.ConfigPath = filepath.Join(homeDir, ".config", "doc", "userConfig.yaml")
	return nil
}

func (u *User) setDefaultNotesLocation() error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working dir: %w", err)
	}
	u.NotesLocation = path
	return nil
}

func (u *User) promptUserName() error {
	fmt.Print("What is your user name? ")
	input, err := readInput()
	if err != nil {
		return fmt.Errorf("failed to readinput: %w", err)
	}
	u.UserName = input
	return nil
}

func (u *User) promptEditor() error {
	fmt.Print("What editor do you want to use by default? (e.g., nvim, code, emacs): ")
	input, err := readInput()
	if err != nil {
		return fmt.Errorf("failed to readinput: %w", err)
	}
	u.Editor = input
	return nil
}

// Helper function to read and clean user input
func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read string: %w", err)
	}
	return strings.TrimSpace(input), nil
}

// TODO: need to dynamically go get config file location and then subsiquent info in it. should marshal it into a user struct. should probably make these methods on User struct.
func GetUserConfig() (*User, error) {
	var user User
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home dir: %w", err)
	}

	p := filepath.Join(homeDir, ".config", "doc", "userConfig.yaml")
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(data, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &user, nil
}
