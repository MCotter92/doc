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
		fmt.Printf("Could not create new user: %s", err)
		return err
	}

	data, err := yaml.Marshal(user)
	if err != nil {
		fmt.Printf("Could not marshal config to YAML: %s", err)
		return err
	}

	configFile := filepath.Join(user.ConfigPath, "userConfig.yaml")

	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		fmt.Printf("Could not write to config file: %s", err)
		return err
	}

	contents, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Could not read config file: %s", err)
		return err
	}

	fmt.Println("Configuration file created at ", user.ConfigPath, ".")
	fmt.Println("========================================")
	fmt.Println("Config:")
	fmt.Println(string(contents))

	database, err := utils.NewDatabase()
	if err != nil {
		fmt.Printf("Could not create database: %s", err)
		return err
	}
	defer database.Close()

	if err := database.CreateTables(); err != nil {
		fmt.Printf("Could not create tables: %s", err)
		return err
	}

	db, err := utils.NewDatabase()
	if err != nil {
		return err
	}
	err = db.InsertUser(user)
	if err != nil {
		return err
	}

	fmt.Printf("Database created at %s\n", database.Path)

	return err
}
