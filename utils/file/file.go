package file

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func moveTo(content os.DirEntry, folderPath string, newPath string) {
    filePath := folderPath + "/" + content.Name()

    os.Rename(filePath, newPath)
}

func getFileExtension(fileName string) (string) {
    for i := len(fileName) - 1; i >= 0; i-- {
        if fileName[i] == '.' {
            return fileName[i+1:]
        }
    }

    return "" 
}
 
func createExtensionDir(file os.DirEntry, path string) (string, error){
    if (file.IsDir()) {
        return "", nil 
    }

    fileName := file.Name()
    extension := getFileExtension(fileName)
    if extension == "" {
        return "", nil
    }

    extensionDir := fmt.Sprintf("%s/%s", path, strings.ToUpper(extension))
    fmt.Println(extensionDir)

    err := os.Mkdir(extensionDir, 0750)
    if err != nil && os.IsExist(err) {
        return "", err
    }

    return extensionDir, nil
}

func OrganizeFiles(folderContent []os.DirEntry, folderPath string) {
    for _, content := range folderContent {
        extensionDir, err := createExtensionDir(content, folderPath)
        if err != nil {
            log.Fatal(err)
        }
        moveTo(content, folderPath, extensionDir)
    }
}

