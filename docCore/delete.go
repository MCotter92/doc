package docCore

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/MCotter92/doc/utils"
)

func Delete(searchRes []utils.Doc, db *utils.Database) error {

	// TODO: move this outside of Delete? so search->TableOutput->CRUD
	utils.TableOutput(searchRes)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("Could not read string: %s", err)
	}

	trimmed := strings.TrimSpace(input)

	numInput, err := strconv.Atoi(trimmed)
	if err != nil {
		return fmt.Errorf("Could not convert string to int: %s", err)
	}

	if numInput < 0 || numInput >= len(searchRes) {
		return fmt.Errorf("Invalid selection: %d. Must be between 0 and %d", numInput, len(searchRes)-1)
	}

	dbInput := searchRes[numInput].Id.String()

	fmt.Printf("DEBUG: id='%s' (len=%d)\n", dbInput, len(dbInput))

	err = db.DeleteDoc(dbInput)
	if err != nil {
		return fmt.Errorf("Could not delete doc from sqlite: %w", err)
	}

	fmt.Print(searchRes[numInput].Path)
	command := exec.Command("rm", searchRes[numInput].Path)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		return fmt.Errorf("Could not run command: %s", err)
	}

	fmt.Println("=================================================================")
	fmt.Println("Document successfully deleted from both database and filesystem.")
	fmt.Println("=================================================================")
	return nil

}
