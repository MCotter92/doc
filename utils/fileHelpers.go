package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func HasNotesInLocation(locatin string) bool {
	if locatin == "" {
		return false
	}

	entries, err := os.ReadDir(locatin)
	if err != nil {
		return false
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			name := entry.Name()
			if filepath.Ext(name) == ".md" {
				return true
			}
		}

	}
	return false

}

func MoveNotes(oldLocation, newLocation string) error {
	if err := os.MkdirAll(newLocation, 0755); err != nil {
		return fmt.Errorf("Failed to create new notes directory: %w", err)
	}

	entries, err := os.ReadDir(oldLocation)
	if err != nil {
		return fmt.Errorf("Failed to read old notes directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			oldFile := filepath.Join(oldLocation, entry.Name())
			newFile := filepath.Join(newLocation, entry.Name())

			if err := os.Rename(oldFile, newFile); err != nil {
				return fmt.Errorf("Failed to move %s: %w", entry.Name(), err)
			}
		}

	}

	return nil
}

func CountNotesInLocation(location string) int {
	if location == "" {
		return -1
	}

	entries, err := os.ReadDir(location)
	if err != nil {
		return -1
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			ext := filepath.Ext(entry.Name())
			if ext == ".md" {
				count++
			}
		}

	}

	return count
}

func IsNoteFile(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".md"

}

func EnsureDirectryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}
