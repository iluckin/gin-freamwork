// This file is part of the TheBears package.
// (c) EnHe <i@iluckin.cn>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package utils

import (
	"github.com/satori/go.uuid"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	PB = 1024 * GB
	EB = 1024 * PB
)

func GenerateRequestId() string {
	return uuid.NewV4().String()
}
