package app

import (
	"harmonybinary/hap"
	"harmonybinary/info"
	"harmonybinary/pkg"
	"harmonybinary/utils"
	"os"
	"path"
)

func OpenFile(filepath string) (*info.App, error) {
	t := utils.GenerateUUid()
	defer os.RemoveAll(t)
	err := utils.Unzip(filepath, t)
	if err != nil {
		return nil, err
	}
	hapFilePath, err := getHapFile(t)
	if err != nil {
		return nil, err
	}
	a, err := hap.OpenFile(hapFilePath)
	return a, nil
}
func getHapFile(dir string) (filePath string, err error) {
	files, err := utils.GetFileList(dir)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		ext := path.Ext(f)
		if ext == pkg.FileTypeHap {
			return f, nil
		}
	}
	return
}
