package info

import (
	"errors"
	"fmt"
	"github.com/shogo82148/androidbinary"
	"github.com/shogo82148/androidbinary/apk"
)

type App struct {
	PackageName  string   `json:"package_name"`
	Name         string   `json:"name"`
	Activities   []string `json:"activities"`
	Version      string   `json:"version"`
	ActivityName string   `json:"activity_name"`
	Perms        []string `json:"perms"`
}

func GetAppInfo(apkFilePath string) (*App, error) {
	p, err := apk.OpenFile(apkFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("打开apk(%s)出错:%s", apkFilePath, err.Error()))
	}
	defer p.Close()
	resConfig := &androidbinary.ResTableConfig{
		Language: [2]uint8{uint8('z'), uint8('h')},
	}
	appLabel, err := p.Label(resConfig)
	if err != nil {
		resConfig = &androidbinary.ResTableConfig{
			Language: [2]uint8{uint8('e'), uint8('n')},
		}
		appLabel, err = p.Label(resConfig)
		if err != nil {
			appLabel = ""
		}
	}
	a := &App{Name: appLabel}
	s, err := p.Manifest().VersionName.String()
	if err == nil {
		a.Version = s
	}
	a.PackageName = p.PackageName()
	activities := p.Manifest().App.Activities
	for _, v := range activities {
		a.Activities = append(a.Activities, v.Name.MustString())
	}
	activity, err := p.MainActivity()
	if err == nil {
		a.ActivityName = activity
	}
	for _, permission := range p.Manifest().UsesPermissions {
		s2, err := permission.Name.String()
		if err == nil {
			a.Perms = append(a.Perms, s2)
		}
	}
	return a, nil
}
