package utils

import "github.com/gin-gonic/gin"

// StandardResponse represents the standard response structure
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CreateStandardResponse creates a standard response with the provided message and data
func Respond(c *gin.Context,r Response)  {
	c.JSON(r.Status, r)
}