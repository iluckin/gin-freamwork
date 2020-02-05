// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package routes

import (
	"github.com/gin-gonic/gin"
)

// define services routes.
func InitializeRouteMap(engine *gin.Engine) *gin.Engine {
	// version v1 routes.
	v1 := engine.Group("v1")
	{
		v1.GET("", func(c *gin.Context) {
			
		})
	}

	return engine
}
