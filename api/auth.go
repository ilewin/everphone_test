package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/transparentideas/everphone_test/utils"
)

type TokenData struct {
	Username string `json:"username"`
}

type ResponseData struct {
	Token        string                 `json:"token"`
	TokenDecoded map[string]interface{} `json:"tokenDecoded"`
}

func (server *Server) GetToken(ctx *gin.Context) {
	token, err := utils.JWTAuthService().Encode(TokenData{Username: "gifty"})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	tokenDecoded, err := utils.JWTAuthService().Decode(token)

	if tokenDecoded.Valid {
		claims := tokenDecoded.Claims.(jwt.MapClaims)

		ctx.JSON(200, gin.H{
			"data": ResponseData{Token: token, TokenDecoded: claims},
			"type": "SUCCESS",
		})
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}

}
