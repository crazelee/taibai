package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"taibai/common/conf"
	"taibai/common/wechat/login"
)

func Login(c *gin.Context) {
	config := conf.Get()
	wxLogin := login.WxConfig{
		AppID:  config.Wechat.Appid,
		Secret: config.Wechat.Secret,
	}

	code := c.Query("code")
	wxUserInfo, err := wxLogin.LoginCode(code)
	if err != nil {
		fmt.Println("code 失效", err)
	}
	fmt.Println(wxUserInfo)

}
