/*
 * // This file is part of the TheBears package.
 * // (c) EnHe <i@iluckin.cn>
 * //
 * // For the full copyright and license information, please view the LICENSE
 * // file that was distributed with this source code.
 */

package services

import "github.com/iluckin/bear/utils"

func Success(data interface{}) *utils.Response {
	return &utils.Response{
		"code":    OK,
		"message": "OK",
		"data":    data,
	}
}

func Failed(message string) *utils.Response {
	return &utils.Response{
		"code":    Fail,
		"message": message,
		"data":    nil,
	}
}

func FailedWithCode(code int) *utils.Response {
	return &utils.Response{
		"code":    code,
		"message": "fail",
		"data":    nil,
	}
}
