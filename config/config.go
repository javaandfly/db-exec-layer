package config

import (
	"db-exec-layer/pkg/utils"

	"github.com/sirupsen/logrus"
)

type DatabaseSource struct {
	DriverName      string `mapstructure:"driver_name"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Database        string `mapstructure:"database"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Charset         string `mapstructure:"charset"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	Debug           bool   `mapstructure:"debug"`
}
type LogConfig struct {
	LogFilePath string `mapstructure:"log_file_path"`
	LogLevel    string `mapstructure:"log_level"`
	LogName     string `mapstructure:"log_name"`
}

type DbLayerConfig struct {
	Log            LogConfig      `mapstructure:"log"`
	DatabaseConfig DatabaseSource `mapstructure:"db"`
}

func ResolveConfig(path string, cfg interface{}) error {
	vp, err := utils.ReadConfig(path)
	if err != nil {
		return err
	}

	//replace according to the given structure
	if err := vp.Unmarshal(&cfg); err != nil {
		logrus.Errorf("unmarshal config file failed %v", err)
		return err
	}
	return nil
}
