package docCore

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MCotter92/doc/utils"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

func DocInit() {

	user, err := utils.NewUser()
	if err != nil {
		fmt.Errorf("Could not create new user: %w", err)
	}
	// Marshal to YAML
	data, err := yaml.Marshal(user)
	if err != nil {
		fmt.Println("Could not marshal config to YAML:", err)
		return
	}

	// Write config file
	err = os.WriteFile(user.ConfigPath, data, 0644)
	if err != nil {
		fmt.Println("Could not write to config file:", err)
		return
	}

	// Read and print the config
	contents, err := os.ReadFile(user.ConfigPath)
	if err != nil {
		fmt.Println("Could not read config file:", err)
		return
	}

	fmt.Println("Configuration file created at ", user.ConfigPath, ".")
	fmt.Println("========================================")
	fmt.Println("Config:")
	fmt.Println(string(contents))

	dbPath := filepath.Join(filepath.Dir(user.ConfigPath), "doc.db")

	database, err := utils.NewDatabase(dbPath)
	if err != nil {
		fmt.Printf("Could not create database: %v\n", err)
		return
	}
	defer database.Close()

	if err := database.CreateTables(); err != nil {
		fmt.Printf("Could not create tables: %v\n", err)
		return
	}

	fmt.Printf("Database created at %s\n", database.Path)

}
