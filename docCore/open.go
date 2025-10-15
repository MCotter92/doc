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

func Open(searchRes []utils.Doc) error {

	utils.TableOutput(searchRes)

	fmt.Println("Select the row numbers of the notes you want to open. Seperate with spaces.")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return fmt.Errorf("Could not read string: %s", err)
	}

	// fmt.Printf("You selected %s\n", input)

	config, err := utils.GetUserConfig()
	if err != nil {
		return fmt.Errorf("Could not get user config: %s", err)
	}

	trimmed := strings.TrimSpace(input)

	numInput, err := strconv.Atoi(trimmed)
	if err != nil {
		return fmt.Errorf("Could not convert string to int: %s", err)
	}

	fmt.Println("selected input: ", searchRes[numInput].Path)
	command := exec.Command(config.Editor, searchRes[numInput].Path)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout

	err = command.Run()
	if err != nil {
		return fmt.Errorf("Could not run command: %s", err)
	}

	return nil
}
