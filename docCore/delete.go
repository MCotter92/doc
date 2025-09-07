package docCore

import "github.com/MCotter92/doc/utils"

func Delete(searchRes []utils.Doc) error {

	// render table for user to choose from
	utils.TableOutput(searchRes)
	// get path from doc struct an delete file
	// get uuid from doc and delete db entry
	return nil

}
