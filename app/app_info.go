package app

import (
	"harmonybinary/hap"
	"harmonybinary/info"
	"harmonybinary/pkg"
	"harmonybinary/utils"
	"io/ioutil"
	"os"
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
		fs, _ := ioutil.ReadFile(f)
		fileType := utils.GetFileType(fs)
		if fileType == pkg.FileTypeHap {
			return f, nil
		}
	}
	return
}
