/*
 * // This file is part of the TheBears package.
 * // (c) EnHe <i@iluckin.cn>
 * //
 * // For the full copyright and license information, please view the LICENSE
 * // file that was distributed with this source code.
 */

package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"os"
)

func ConfigureLogger() gin.LoggerConfig {
	return gin.LoggerConfig{
		Formatter: loggerFormatter,
		Output:    loggerOutputs(),
	}
}

func loggerOutputs() io.Writer {
	if file, err := os.Create(viper.GetString("server.log.accept")); err == nil {
		return file
	}

	return nil
}

// logger lime formatter.
func loggerFormatter(p gin.LogFormatterParams) string {
	template := "[%s-%s] | %s | %d | %s | %s | %s | %s | %s | %s | %s\n"

	return fmt.Sprintf(template,
		viper.GetString("server.name"),
		p.Request.Header.Get("X-Request-Id"),
		p.TimeStamp.Format("2006-01-02 15:04:05"),
		p.StatusCode,
		p.Latency,
		p.ClientIP,
		p.Method,
		p.Path,
		p.Request.Proto,
		p.ErrorMessage,
		p.Request.UserAgent(),
	)
}
