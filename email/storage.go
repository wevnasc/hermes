package email

import (
	"fmt"
	"os"
)

type storage interface {
	exists(file string) bool
}

type localStorate struct {
	rootFolder string
}

func NewLocalStorate(rootFolder string) *localStorate {
	return &localStorate{rootFolder: rootFolder}
}

func (ls *localStorate) exists(file string) bool {
	filePath := fmt.Sprintf("%s/%s", ls.rootFolder, file)
	if _, err := os.Stat(filePath); err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
