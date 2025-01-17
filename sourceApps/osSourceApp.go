package sourceapps

import (
	"github.com/probeldev/fastlauncher/model"
	"github.com/probeldev/fastlauncher/pkg/finderallapps"
)

type OsSourceApp struct{}

func (o *OsSourceApp) GetAll() ([]model.App, error) {
	os := o.getOs()
	finder, err := finderallapps.GetFinder(os)

	if err != nil {
		return nil, err
	}

	osApps, err := finder.GetAllApp()
	if err != nil {
		return nil, err
	}

	apps := make([]model.App, len(osApps))

	for _, oa := range osApps {
		apps = append(apps, model.App{
			Title:       oa.Name,
			Description: oa.Description,
			Command:     oa.Command,
			Keywords:    oa.Keywords,
		})
	}

	return apps, nil

}

func (o *OsSourceApp) getOs() string {
	// TODO change
	return finderallapps.OsLinux
}