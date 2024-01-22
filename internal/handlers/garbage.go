package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGarbage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, `{"msg": "hello"}`)
}
