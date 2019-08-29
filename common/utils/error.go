package utils

import "taibai/consts"

func GetErrMsg(code int) string {
	return consts.ErrMsg[code]
}

