package storage

import (
	"fmt"
	"os"
)

type local struct {
	rootFolder string
}

func NewLocal(rootFolder string) Storager {
	return &local{rootFolder: rootFolder}
}

func (ls *local) Exists(file string) bool {
	filePath := fmt.Sprintf("%s/%s", ls.rootFolder, file)
	if _, err := os.Stat(filePath); err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
