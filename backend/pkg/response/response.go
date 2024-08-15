package response

import "github.com/gin-gonic/gin"

type R struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
}

type Response struct {
}

func Success(c *gin.Context, resp R) {
	resp.Status = 0

	c.JSON(200, resp)
}

func Ok(c *gin.Context, msg string) {
	Success(c, R{
		Msg: msg,
	})
}

func Fail(c *gin.Context, resp R) {
	resp.Status = 1

	c.JSON(200, resp)
}
