package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// read configuration file
func ReadConfig(filePath string) (*viper.Viper, error) {
	vp := viper.New()
	// set path
	vp.SetConfigFile(filePath)

	// compatible with toml type because .conf and .cxt cannot be parsed
	fileType := filepath.Ext(filePath)
	if fileType == ".conf" || fileType == ".cfg" {
		vp.SetConfigType("toml")
	}

	vp.AutomaticEnv()

	//the . in the field is replaced by _
	replacer := strings.NewReplacer(".", "_")
	vp.SetEnvKeyReplacer(replacer)

	// read file
	if err := vp.ReadInConfig(); err != nil {
		err = fmt.Errorf("read config file failed %v", err)
		return nil, err
	}

	return vp, nil

}
