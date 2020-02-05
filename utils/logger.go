// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package utils

import (
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func NewLogger() error {
	cfg := log.PassLagerCfg{
		Writers:        viper.GetString("server.log.channel"),
		LoggerLevel:    viper.GetString("server.log.level"),
		LoggerFile:     viper.GetString("server.log.path"),
		LogFormatText:  viper.GetString("server.log.format") == "text",
		RollingPolicy:  viper.GetString("server.log.rolling"),
		LogRotateDate:  viper.GetInt("server.log.rotate.date"),
		LogRotateSize:  viper.GetInt("server.log.rotate.size"),
		LogBackupCount: viper.GetInt("server.log.backup"),
	}

	log.InitWithConfig(&cfg)

	return nil
}
