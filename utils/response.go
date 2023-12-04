package util

import "github.com/gin-gonic/gin"

type response struct {
	Data interface{} `json:"data,omitempty"`
}

func APIResponse(c *gin.Context, StatusCode int, Data interface{}) {
	response := response{
		Data: Data,
	}

	if StatusCode >= 400 {
		c.JSON(StatusCode, response)
		defer c.AbortWithStatus(StatusCode)
	} else {
		c.JSON(StatusCode, response)
	}
}
