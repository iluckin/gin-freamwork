// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iluckin/bear/handlers/health"
	"github.com/lexkong/log"
)

// define health check api
func InitializeHealthRouteMap(engine *gin.Engine) *gin.Engine {
	log.Info("Define health-check routes...")

	healthApi := engine.Group("/v1/health")
	{
		healthApi.GET("", health.Health)
		healthApi.GET("/cpu", health.Cpu)
		healthApi.GET("/disk", health.Disk)
		healthApi.GET("/mem", health.Memory)
	}

	return engine
}
