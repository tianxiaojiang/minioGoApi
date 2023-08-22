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
		// 如果头信息里没有，取url的参数
		if sign == "" {
			sign, _ = c.GetQuery("sign")
		}

		// 检查timestamp是否存在
		timeStr := c.GetHeader("timestamp")
		if timeStr == "" {
			timeStr, _ = c.GetQuery("timestamp")
		}

		timeInt, err := strconv.ParseInt(timeStr, 10, 64)
		if err != nil {
			c.Data(403, "plain/text", []byte("timestamp error."))
			c.Abort()
			return
		}
		timeStampT := time.Unix(timeInt, 0)

		// 时间戳与服务端相差不超过10s
		// 暂时不校验时间，方便缓存
		if t.Sub(timeStampT).Seconds() > 15 || timeStampT.Sub(t).Seconds() > 15 {
			// c.Data(403, "plain/text", []byte("timestamp error."))
			// c.Abort()
			// return
			err = nil
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
