package config

import "github.com/spf13/viper"

// Will contain configuration of the application,
// the values are read by viper from a config file or env.

func LoadConfig(path string, name string, ext string, cfg interface{}) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(ext) // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(cfg)
	return
}
