package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type User struct {
	ID             uuid.UUID `mapstructure:"id" yaml:"id"`
	UserName       string    `mapstructure:"userName" yaml:"userName"`
	NotesLocation  string    `mapstructure:"notesLocation" yaml:"notesLocation"`
	Editor         string    `mapstructure:"editor" yaml:"editor"`
	ConfigDir      string    `mapstructure:"editor" yaml:"ConfigDir"`
	ConfigFilePath string    `mapstructure:"configPath" yaml:"configFilePath"`
	DbPath         string    `mapstructure:"configPath" yaml:"dbPath"`
}

// Top level function for user creation. Fills out the user's struct.
func NewUser() (*User, error) {
	user := &User{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to set config path: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "doc")
	user.ConfigDir = configDir

	configFilePath := filepath.Join(homeDir, ".config", "doc", "configFile.yaml")
	user.ConfigFilePath = configFilePath

	dbPath := filepath.Join(homeDir, ".config", "doc", "db.sql")
	user.DbPath = dbPath

	viper.SetDefault("ConfigDir", configDir)
	viper.SetDefault("dbPath", dbPath)
	viper.SetDefault("configFilePath", configFilePath)
	viper.SetDefault("editor", "nvim")

	if err := setupViper(); err != nil {
		return nil, fmt.Errorf("Failed to setup viper: %w", err)
	}

	user.setUserID()
	user.promptUserName()

	if err := user.promptNotesLocation(); err != nil {
		return nil, fmt.Errorf("failed to set notes location: %w", err)
	}
	user.makeNotesLocation()

	if err := user.promptEditor(); err != nil {
		return nil, fmt.Errorf("failed to set editor: %w", err)
	}

	if err := user.makeConfigLocation(user.ConfigDir, user.ConfigFilePath); err != nil {
		return nil, fmt.Errorf("Failed to make config location: %w", err)
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
	viper.Set("configDir", u.ConfigDir)
	viper.Set("configFilePath", u.ConfigFilePath)
	viper.Set("dbPath", u.DbPath)

	if err := viper.WriteConfigAs(u.ConfigFilePath); err != nil {
		return fmt.Errorf("Could not write to config file: %w", err)
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
	u.ID = uuid.New()
}

func (u *User) promptNotesLocation() error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Could not get pwd: %w", err)
	}
	fmt.Println("Would you like to set your notes directory at the current location: %s", pwd)
	fmt.Println("Please enter y or n.")

	var response string
	fmt.Scanln(&response)
	if response == "y" || response == "Y" {
		u.NotesLocation = pwd
		return nil
	} else if response == "n" || response == "N" {
		var location string
		fmt.Println("Please provide an alternate notes directory.")
		fmt.Scan(&location)
		u.NotesLocation = location
	}

	return nil

}

func (u *User) makeNotesLocation() error {
	notesLocation := filepath.Join(u.NotesLocation, "notes")
	if err := os.MkdirAll(notesLocation, 0755); err != nil {
		return fmt.Errorf("Failed to make Notes Location: %w", err)
	}

	return nil

}

func (u *User) makeConfigLocation(configDir, configFilePath string) error {

	if err := os.MkdirAll(configDir, 0775); err != nil {
		return fmt.Errorf("Failed to create config dir: %w", err)
	}

	if _, err := os.Create(configFilePath); err != nil {
		return fmt.Errorf("Failed to create config file: %w", err)
	}

	return nil
}

func (u *User) promptUserName() {
	fmt.Println("Please provide a user name.")
	var userName string
	fmt.Scanln(&userName)
	u.UserName = userName
}

func (u *User) promptEditor() error {
	fmt.Print("What editor do you want to use by default? (e.g., nvim, code, emacs): ")
	var input string
	fmt.Scanln(&input)

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
