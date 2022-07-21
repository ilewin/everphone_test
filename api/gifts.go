package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/transparentideas/everphone_test/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *Server) ListGifts(c *gin.Context) {
	gifs, err := server.store.ListGifts(c.Request.Context(), db.ListGiftsParams{})
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, gifs)
}

func (server *Server) GetGift(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}

	gif, err := server.store.GetGift(c.Request.Context(), db.GetGiftParams{Id: id})
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, gif)
}

func (server *Server) AddGift(c *gin.Context) {
	var req db.AddGiftParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	gif, err := server.store.AddGift(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, gif)
}

func (server *Server) UpdateGift(c *gin.Context) {
	var req db.UpdateGiftParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	gift, err := server.store.UpdateGift(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, gift)
}
