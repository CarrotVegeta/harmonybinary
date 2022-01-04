package hap

import (
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
	apkFilePath, err := getApkFile(t)
	if err != nil {
		return nil, err
	}
	a, err := info.GetAppInfo(apkFilePath)
	return a, nil
}
func getApkFile(dir string) (filePath string, err error) {
	files, err := utils.GetFileList(dir)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		f, _ := ioutil.ReadFile(f)
		fileType := utils.GetFileType(f)
		if fileType == pkg.FileTypeApk {
			return fileType, nil
		}
	}
	return
}
