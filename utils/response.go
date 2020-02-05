/*
 * // This file is part of the TheBears package.
 * // (c) EnHe <i@iluckin.cn>
 * //
 * // For the full copyright and license information, please view the LICENSE
 * // file that was distributed with this source code.
 */

package utils

import "github.com/gin-gonic/gin"

type Response gin.H

type ResponseInterface interface {
	SetCode(code int) *Response
	SetData(data interface{}) *Response
	SetMessage(message string) *Response
}
