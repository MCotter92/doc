package docCore

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MCotter92/doc/utils"
)

type ConfigUpdateRequest struct {
	UserName      string
	NotesLocation string
	Editor        string
	ConfigPath    string
}

// updates both users table and config file
func UpdateUserConfiguration(req ConfigUpdateRequest) error {

	user, err := utils.GetUserConfig()
	if err != nil {
		return fmt.Errorf("Faild to get user config: %w", err)
	}

	fmt.Println("Updating user config...")
	var updated bool

	if req.UserName != "" {
		oldUserName := user.UserName
		if err = user.UpdateConfigFile("UserName", req.UserName); err != nil {
			return fmt.Errorf("Failed to update notes location: %w", err)
		}

		fmt.Printf("Notes location updated: %s -> %s\n", oldUserName, req.UserName)

	}

	if req.Editor != "" {
		oldEditor := user.Editor
		if err = user.UpdateConfigFile("editor", req.Editor); err != nil {
			return fmt.Errorf("Failed to update editor: %w", err)
		}

		fmt.Printf("Editor udpated: %s -> %s\n", oldEditor, req.Editor)
		updated = true
	}

	if req.NotesLocation != "" {
		if err := handleNotesLocationUpdate(user, req.NotesLocation); err != nil {
			return fmt.Errorf("Failed to update notes location: %w", err)
		}
		updated = true
	}

	if updated {
		fmt.Println("\nConfiuration updated sucessfully!")

		if err := user.Validate(); err != nil {
			fmt.Errorf("Configuragion validation failed: %w", err)
		}

		fmt.Println("\nUpdated Configuragion!")
		// call UpdateUserTable() here?
		return ShowUserConfig(false)
	}

	return nil

}

// func UpdateUserTable(req ConfigUpdateRequest) error {}

func ShowUserConfig(showDetailed bool) error {
	user, err := utils.GetUserConfig()
	if err != nil {
		return fmt.Errorf("Failed to load config: %w", err)
	}

	fmt.Println("Current Config")
	fmt.Println("==========================================")
	fmt.Printf("Username:       %s\n", user.UserName)
	fmt.Printf("Editor:         %s\n", user.Editor)
	fmt.Printf("Notes Location: %s\n", user.NotesLocation)
	fmt.Printf("User ID:        %s\n", user.ID)
	fmt.Println("==========================================")

	if showDetailed {
		fmt.Printf("Config Path:       %s\n", user.ConfigPath)
	}

	if noteCount := utils.CountNotesInLocation(user.NotesLocation); noteCount >= 0 {
		fmt.Printf("Note Count:       %d\n", noteCount)
	}

	return nil
}

func handleNotesLocationUpdate(user *utils.User, newLocation string) error {
	oldLocation := user.NotesLocation

	if newLocation[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("Failed to get home directory: %w", err)
		}
		newLocation = filepath.Join(homeDir, newLocation[1:])
	}

	absPath, err := filepath.Abs(newLocation)
	if err != nil {
		return fmt.Errorf("Failed to resolve path: %w", err)
	}

	if oldLocation != "" && oldLocation != absPath {
		if utils.HasNotesInLocation(oldLocation) {
			fmt.Printf("Notes found in old location: %s\n", oldLocation)
			fmt.Printf("New location will be: %s\n", absPath)
			fmt.Print("Do you want to move existing notes to the new location? [y/N]: ")

			var response string
			fmt.Scanln(&response)

			if response == "y" || response == "Y" || response == "yes" {
				if err := utils.MoveNotes(oldLocation, absPath); err != nil {
					return fmt.Errorf("failed to move notes: %w", err)
				}
				fmt.Println("Notes moved successfully!")
			} else {
				fmt.Println(" Notes were not moved. You can manually move them later.")
			}
		}
	}

	// Update the configuration
	if err := user.UpdateConfigFile("notesLocation", absPath); err != nil {
		return fmt.Errorf("Failed to update notes location: %w", err)
	}

	fmt.Printf("Notes location updated: %s -> %s\n", oldLocation, absPath)
	return nil
}
