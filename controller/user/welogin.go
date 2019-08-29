package user

import (
	"encoding/json"
	"fmt"
	"io"
	"taibai/common/log"
	"taibai/common/utils"
	"taibai/consts"
	"taibai/middleware"

	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp"
	"taibai/common/conf"
)

type WeLoginReq struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}

type WeLoginResp struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func WeLogin(c *gin.Context) {
	var req WeLoginReq
	var errNo int

	defer func() {
		c.JSON(200, gin.H{
			"error": errNo,
			"errmsg": utils.GetErrMsg(errNo),
			"data": map[string]interface{}{
				"code": req.Code,
			},
		})
	}()

	// 获取参数
	decoder := json.NewDecoder(c.Keys[middleware.TotalParamsKey].(io.Reader))
	if err := decoder.Decode(&req); err != nil {
		log.Trace.Error(c, "params illegal", err)
		errNo = consts.ParamsDecodeError
		return
	}

	config := conf.Get()
	res, err := weapp.Login(config.Wechat.Appid, config.Wechat.Secret, req.Code)
	if err != nil {
		// 处理一般错误信息
		log.Trace.Errorf("code: %v, err : %v", req, err)
		errNo = consts.NetWorkError
		return
	}

	if res.HasError() {
		// 处理微信返回错误信息
		log.Trace.Errorf("code: %v, err : %v", req, err)
		errNo = consts.NetWorkError
		return
	}
	res.HasError()

	fmt.Printf("返回结果: %#v", res)
}
