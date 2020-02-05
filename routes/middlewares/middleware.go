// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/iluckin/bear/utils"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Headers map[string]string

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.

var noCacheHeaders = Headers{
	"Expires":       "Thu, 01 Jan 1970 00:00:00 GMT",
	"Last-Modified": time.Now().UTC().Format(http.TimeFormat),
	"Cache-Control": "no-cache, no-store, max-age=0, must-revalidate, value",
}

func NoCache(c *gin.Context) {
	for k, v := range noCacheHeaders {
		c.Header(k, v)
	}

	c.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.

var allowOptionRequestHeaders = Headers{
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Methods": "GET,POST,PUT,PATCH,DELETE,OPTIONS",
	"Access-Control-Allow-Headers": "authorization, origin, content-type, accept",
	"Allow":                        "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS",
}

func AllowOptionRequest(c *gin.Context) {
	if c.Request.Method != http.MethodOptions {
		c.Next()
	} else {
		for k, v := range allowOptionRequestHeaders {
			c.Header(k, v)
		}
	}
}

// Accept application/json
func AcceptJson(c *gin.Context) {
	c.Header("Content-Type", "application/json")
}

var secureHeaders = Headers{
	"Access-Control-Allow-Origin": "*",
	"X-Frame-Options":             "DENY",
	"X-Content-Type-Options":      "nosniff",
	"X-XSS-Protection":            "1; mode=block",
}

// Also consider adding Content-Security-Policy headers
// c.Headers("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")

func Secure(c *gin.Context) {
	for k, v := range secureHeaders {
		c.Header(k, v)
	}

	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}

// custom server flag.
func Server(c *gin.Context) {
	requestId := utils.GenerateRequestId()
	c.Request.Header.Set("X-Request-Id", requestId)

	c.Header("X-Request-Id", requestId)
	c.Header("Server", viper.GetString("server.flag"))
}
