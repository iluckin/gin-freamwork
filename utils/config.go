// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package utils

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	FileName string
}

// initialize config with given config file dir.
func (c *Config) initialize() error {
	if c.FileName != "" {
		viper.SetConfigFile(c.FileName)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("cfg")
	}

	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TB_")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// watch a config file changed.
func (c *Config) watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("config file changed: %s", e.Name)
	})
}

// get config instance.
func NewConfig(file string) error {
	c := Config{FileName: file}
	if err := c.initialize(); err != nil {
		return err
	}

	c.watch()

	return nil
}
