package usecases

import (
	"encoding/base64"
)

func (i interactor) getKubeconfigFromSettings() string {
	config := i.settingsRepo.GetSettingsValue("kubeconfig")

	if config == "" {
		return ""
	}

	sDec, err := base64.StdEncoding.DecodeString(config)

	if err != nil {
		return ""
	}

	return string(sDec)
}
