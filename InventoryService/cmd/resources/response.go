package resources

import "github.com/gin-gonic/gin"

type DataResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	ErrMsg string `json:"message"`
}

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, httpCode int, msg string, data interface{}) {
	if msg != "" {
		if data != nil {
			c.JSON(httpCode, Response{
				Msg:  msg,
				Data: data,
			})
			return
		}
		c.JSON(httpCode, ErrorResponse{
			ErrMsg: msg,
		})
		return
	}
	c.JSON(httpCode, DataResponse{
		Data: data,
	})
}
