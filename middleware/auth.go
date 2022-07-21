package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/transparentideas/everphone_test/utils"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type invalidAuthHeader struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, invalidAuthHeader{
				Field: "Authorization",
				Value: "",
				Tag:   "required",
				Param: "",
			})
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) != 2 {
			c.JSON(500, gin.H{
				"code": "TOKEN_NOT_FOUND",
				"type": "ERROR",
			})
			c.Abort()
			return
		}

		_, err := utils.JWTAuthService().Decode(idTokenHeader[1])
		if err != nil {
			c.JSON(500, gin.H{
				"code": "TOKEN_INVALID",
				"type": "ERROR",
			})
			c.Abort()
			return
		}
		c.Next()

	}
}
