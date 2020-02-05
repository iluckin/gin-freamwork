// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"github.com/iluckin/bear/bootstrap"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cf = pflag.StringP("config", "c", "cfg.toml", "the config file path.")
)

func main() {
	pflag.Parse()

	app := bootstrap.Initialize(cf)
	listen := viper.GetString("server.listen")

	log.Infof("listening on http address: http://%s", listen)
	log.Infof(app.Run(listen).Error())
}
