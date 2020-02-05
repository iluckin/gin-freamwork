// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/iluckin/bear/routes"
	"github.com/iluckin/bear/utils"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"os"
)

func Initialize(config *string) *gin.Engine {

	if err := utils.NewConfig(*config); err != nil {
		log.Fatal("Config component initialization failed.", err)
	}

	if err := utils.NewLogger(); err != nil {
		log.Fatal("Logger component initialization failed.", err)
	}

	acceptLogFile, err := os.Create(viper.GetString("server.log.accept"))
	if err != nil {
		log.Fatal("", err)
	}

	log.Info("Starting service...")
	gin.SetMode(viper.GetString("server.env"))
	gin.DisableConsoleColor()
	gin.DefaultWriter = acceptLogFile

	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(ConfigureLogger()))
	engine.Use(gin.Recovery())
	engine.NoRoute(routes.HandleNoRoute)

	routes.LoadDefaultMiddleware(engine)
	routes.InitializeHealthRouteMap(engine)
	routes.InitializeRouteMap(engine)

	go utils.SelfCheck(viper.GetString("server.health.api"))

	return engine
}
