package hap

import (
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
	apkFilePath, err := getApkFile(t)
	if err != nil {
		return nil, err
	}
	a, err := info.GetAppInfo(apkFilePath)
	if err != nil {
		return nil, err
	}
	return a, nil
}
func getApkFile(dir string) (filePath string, err error) {
	files, err := utils.GetFileList(dir)
	if err != nil {
		return "", err
	}
	for _, f := range files {
		ext := path.Ext(f)
		if ext == pkg.FileTypeApk {
			return f, nil
		}
	}
	return
}
