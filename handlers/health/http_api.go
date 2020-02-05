// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package health

import (
	"github.com/gin-gonic/gin"
	"github.com/iluckin/bear/handlers"
	"github.com/iluckin/bear/services"
	"github.com/lexkong/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"net/http"
	"strconv"
)

// HealthCheck shows `OK` as the ping-pong result.
func Health(c *gin.Context) {
	log.Infof("Client[%s] accept health api.", c.ClientIP())

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"code":    handlers.OK,
	})
}

// checks the disk usage.
func Disk(c *gin.Context) {
	stat, e := disk.Usage("/")

	if e != nil {
		c.JSON(http.StatusOK, services.Failed("Failed to get disk information."))
	}

	percent := int(stat.UsedPercent)
	message := "GREEN"
	// TODO If you need to optimize here, for now this is just an example
	switch {
	case percent >= 90:
		message = "YELLOW"
	case percent >= 95:
		message = "RED"
	}

	data := map[string]interface{}{
		"disk_usage_percent": strconv.Itoa(percent) + "%",
		"disk_usage_level":   message,
	}

	c.JSON(http.StatusOK, services.Success(data))
}

// checks the cpu usage.
func Cpu(c *gin.Context) {
	cores, e := cpu.Counts(false)

	if e != nil {
		c.JSON(http.StatusOK, services.FailedWithCode(services.InternalError))
	}

	data := gin.H{
		"cores": cores,
	}

	c.JSON(http.StatusOK, services.Success(data))
}

// checks the memory usage.
func Memory(c *gin.Context) {
	stat, e := mem.VirtualMemory()

	if e != nil {
		c.JSON(http.StatusOK, services.Failed("Failed to get memory information."))
	}

	c.JSON(http.StatusOK, services.Success(stat.Available))
}
