package docCore

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MCotter92/doc/utils"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

func DocInit() error {

	user, err := utils.NewUser()
	if err != nil {
		fmt.Printf("Could not create new user: %w", err)
		return err
	}
	// Marshal to YAML
	data, err := yaml.Marshal(user)
	if err != nil {
		fmt.Printf("Could not marshal config to YAML: %w", err)
		return err
	}

	// Write config file
	err = os.WriteFile(user.ConfigPath, data, 0644)
	if err != nil {
		fmt.Printf("Could not write to config file: %w", err)
		return err
	}

	// Read and print the config
	contents, err := os.ReadFile(user.ConfigPath)
	if err != nil {
		fmt.Printf("Could not read config file: %w", err)
		return err
	}

	fmt.Println("Configuration file created at ", user.ConfigPath, ".")
	fmt.Println("========================================")
	fmt.Println("Config:")
	fmt.Println(string(contents))

	dbPath := filepath.Join(filepath.Dir(user.ConfigPath), "doc.db")

	database, err := utils.NewDatabase(dbPath)
	if err != nil {
		fmt.Printf("Could not create database: %w", err)
		return err
	}
	defer database.Close()

	if err := database.CreateTables(); err != nil {
		fmt.Printf("Could not create tables: %w", err)
		return err
	}

	fmt.Printf("Database created at %s\n", database.Path)

	return err
}
