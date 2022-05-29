package ssocreds

import (
	"path/filepath"

	"github.com/bigkevmcd/go-configparser"
	"github.com/winebarrel/ssocreds/utils"
)

const (
	awsConfigPath = ".aws/config"
)

func awsConfig() (*configparser.ConfigParser, error) {
	path := filepath.Join(utils.HomeDir(), awsConfigPath)
	return configparser.NewConfigParserFromFile(path)
}

func SsoStartUrl(profile string) (string, error) {
	config, err := awsConfig()

	if err != nil {
		return "", err
	}

	return config.Get("profile "+profile, "sso_start_url")
}
