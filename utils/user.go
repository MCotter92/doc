package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type User struct {
	ID            string `mapstructure:"id" yaml:"id"`
	UserName      string `mapstructure:"userName" yaml:"userName"`
	NotesLocation string `mapstructure:"notesLocation" yaml:"notesLocation"`
	Editor        string `mapstructure:"editor" yaml:"editor"`
	ConfigPath    string `mapstructure:"configPath" yaml:"configPath"`
}

func NewUser() (*User, error) {
	user := &User{}

	if err := setupViper(); err != nil {
		return nil, fmt.Errorf("Failed to setup viper: %w", err)
	}
	user.setUserID()

	if err := user.setDefaultConfigPath(); err != nil {
		return nil, fmt.Errorf("failed to set config path: %w", err)
	}

	if err := user.setDefaultNotesLocation(); err != nil {
		return nil, fmt.Errorf("failed to set notes location: %w", err)
	}

	viper.SetDefault("id", user.ID)
	viper.SetDefault("configPath", user.ConfigPath)
	viper.SetDefault("notesLocation", user.NotesLocation)
	viper.SetDefault("editor", "nvim")

	if err := user.promptUserName(); err != nil {
		return nil, fmt.Errorf("failed to set user name: %w", err)
	}

	if err := user.promptEditor(); err != nil {
		return nil, fmt.Errorf("failed to set editor: %w", err)
	}

	if err := user.saveConfigFile(); err != nil {
		return nil, fmt.Errorf("Failed to save config: %w", err)
	}

	return user, nil
}

func (u *User) saveConfigFile() error {
	viper.Set("id", u.ID)
	viper.Set("userName", u.UserName)
	viper.Set("notesLocation", u.NotesLocation)
	viper.Set("editor", u.Editor)
	viper.Set("configPath", u.ConfigPath)

	configDir := filepath.Dir(u.ConfigPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("Failed to write confit directory: %w", err)
	}

	if err := viper.WriteConfigAs(u.ConfigPath); err != nil {
		return fmt.Errorf("Failed to write config file: %w", err)
	}

	return nil
}

func (u *User) UpdateConfigFile(key string, value interface{}) error {
	viper.Set(key, value)

	switch key {
	case "userName":
		u.UserName = fmt.Sprintf("%v", value)
	case "editor":
		u.Editor = fmt.Sprintf("%v", value)
	case "notesLocation":
		u.NotesLocation = fmt.Sprintf("%v", value)
	}

	return u.saveConfigFile()
}

func (u *User) Validate() error {
	if u.UserName == "" {
		return fmt.Errorf("userName cannot be empty")
	}
	if u.Editor == "" {
		return fmt.Errorf("editor cannot be empty")
	}
	if u.NotesLocation == "" {
		return fmt.Errorf("notesLocation cannot be empty")
	}

	if _, err := os.Stat(u.NotesLocation); os.IsNotExist(err) {
		if err := os.MkdirAll(u.NotesLocation, 0755); err != nil {
			return fmt.Errorf("failed to create notes directory: %w", err)
		}
	}

	return nil
}

func (u *User) setUserID() {
	u.ID = uuid.New().String()
}

func (u *User) setDefaultConfigPath() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to set config path: %w", err)
	}
	u.ConfigPath = filepath.Join(homeDir, "Documents", "Notes")
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
	if strings.TrimSpace(input) == "" {
		// Use default if empty
		u.Editor = "nvim"
	} else {
		u.Editor = input
	}
	return nil
}

func GetUserConfig() (*User, error) {
	if err := setupViper(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found - please run initializatin first.")
		}
		return nil, fmt.Errorf("cound not read config: %w", err)

	}
	var user User
	if err := viper.Unmarshal(&user); err != nil {
		return nil, fmt.Errorf("could not unmarshl config: %w", err)
	}

	return &user, nil
}

func setupViper() error {
	viper.SetConfigName("userConfig")
	viper.SetConfigType("yaml")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Failed to get home directory: %w", err)
	}

	confgiDir := filepath.Join(homeDir, ".config", "doc")
	viper.AddConfigPath(confgiDir)

	viper.SetEnvPrefix("DOC")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}

func ConfigExists() bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false
	}

	configPath := filepath.Join(homeDir, ".config", "doc", "userConfig.yaml")
	_, err = os.Stat(configPath)
	return err == nil
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read string: %w", err)
	}
	return strings.TrimSpace(input), nil
}
