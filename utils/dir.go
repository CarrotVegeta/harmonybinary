package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetFileList(path string) (files []string, err error) {
	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("filepath.Walk() returned %v\n", err))
	}
	return
}
