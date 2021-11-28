package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{})  {
	c.JSON(status, body)
	return
}

func RespondError(c *gin.Context, err *ApplicationError)  {
	c.JSON(err.StatusCode, err)
	return
}
