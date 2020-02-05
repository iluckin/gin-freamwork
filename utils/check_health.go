// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package utils

import (
	"errors"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func SelfCheck(healthApi string) error {
	if viper.GetBool("server.health.sw") == false {
		return nil
	}

	for i := 0; i < viper.GetInt("server.health.retry"); i++ {
		resp, err := http.Get(healthApi)

		if err == nil && resp.StatusCode == http.StatusOK {
			log.Info("Service started successfully.")

			return nil
		}

		log.Debugf("self-check...")
		time.Sleep(500 * time.Millisecond)
	}

	// send warning message.
	message := "The application has no response, or it might took too long to start up."
	log.Fatal(message, nil)

	return errors.New(message)
}
