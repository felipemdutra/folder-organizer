package folder

import (
    "log"
    "os"
)

func GetFolderContent(folderPath string) ([]os.DirEntry, error) {
    folderContent, err := os.ReadDir(folderPath)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    return folderContent, nil
}

