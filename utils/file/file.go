package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func moveTo(content os.DirEntry, folderPath string, newPath string) error {
    if content.IsDir() {
        return nil
    }

    filePath := folderPath + "/" + content.Name()

    if err := os.Rename(filePath, newPath); err != nil {
        return err
    }

    return nil
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

    if _, err := os.Stat(extensionDir); err == nil {
        return extensionDir, nil
    } else if !os.IsNotExist(err) {
        return "", err
    }

    err := os.Mkdir(extensionDir, 0750)
    if err != nil && os.IsNotExist(err) {
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

        var targetDir string
        
        if extensionDir == "" {
            miscDir := filepath.Join(folderPath, "MISC")
            if _, err := os.Stat(miscDir); os.IsNotExist(err) {
                if err := os.Mkdir(miscDir, 0750); err != nil {
                    log.Fatal(err)
                }
            }
            targetDir = miscDir
        } else {
            targetDir = extensionDir
        }

        newFilePath := filepath.Join(targetDir, content.Name())
        if err := moveTo(content, folderPath, newFilePath); err != nil {
            fmt.Printf("Failed to move file %s: %v\n", content.Name(), err) 
        }
    }
}

