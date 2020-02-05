// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iluckin/bear/routes/middlewares"
	"github.com/lexkong/log"
	"net/http"
)

func HandleNoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "The incorrect API route.",
	})

	return
}

// load default middleware
func LoadDefaultMiddleware(engine *gin.Engine) {
	log.Info("Load defaults middleware...")

	engine.Use(middlewares.Server)
	engine.Use(middlewares.Secure)
	engine.Use(middlewares.NoCache)
	engine.Use(middlewares.AcceptJson)
	engine.Use(middlewares.AllowOptionRequest)
}
