package main

import "github.com/gin-gonic/gin"


type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}



func Respond(c *gin.Context,r Response)  {
	c.JSON(r.Status, r)
}