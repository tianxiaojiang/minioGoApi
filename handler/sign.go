package handler

import (
	"crypto/md5"
	"fmt"
	"storeObj/global"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// get header sign
		sign := c.GetHeader("sign")

		// 检查timestamp是否存在
		timeStr := c.GetHeader("timestamp")
		timeInt, err := strconv.ParseInt(timeStr, 10, 64)
		if err != nil {
			c.Data(403, "plain/text", []byte("timestamp error"))
			c.Abort()
			return
		}
		timeStampT := time.Unix(timeInt, 0)

		// 时间戳与服务端相差不超过10s
		if t.Sub(timeStampT).Seconds() > 10 || timeStampT.Sub(t).Seconds() > 10 {
			c.Data(403, "plain/text", []byte("timestamp error"))
			c.Abort()
			return
		}

		// 检查签名是否正确
		// 签名规则 md5(global.Config.App.SignKey + timestamp)
		s := global.Conf.App.SignKey + timeStr
		serverSign := fmt.Sprintf("%x", md5.Sum([]byte(s)))
		if sign != serverSign {
			c.Data(403, "plain/text", []byte("sign error"))
			c.Abort()
			return
		}

		c.Next()
	}
}
