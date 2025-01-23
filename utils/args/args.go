package args

import (
	"folder-organizer/utils/file"
	"folder-organizer/utils/folder"
	"os"
)

func OrganizeFolder() {
    folderPath := os.Args[len(os.Args)-1]

    content, err := folder.GetFolderContent(folderPath)
    if err != nil {
        return
    }

    file.OrganizeFiles(content, folderPath)
}

