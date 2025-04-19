package docCore

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MCotter92/doc/utils"
)

func CreateProfile(editor, NotesLocation string) {

	user := utils.Profile{}
	user.NewProfile(editor, NotesLocation)

	data, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		fmt.Println("Could not marshal json: ", err)
	}

	err = os.WriteFile("/Users/mason_cotter/dev/doc/utils/profile.json", data, 0644)
	if err != nil {
		fmt.Println("Could not write to profile.json: ", err)
	}

}
