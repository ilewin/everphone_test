package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/transparentideas/everphone_test/db"
)

func (server *Server) MatchGiftToEmployee(c *gin.Context) {
	var req db.MatchGiftToEmployeeParams
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	emplsgifts, err := server.store.MatchGiftToEmployee(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	c.JSON(200, emplsgifts)
}

func (server *Server) MatchAll(c *gin.Context) {
	var req db.MatchAllParams
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	emplsgifts, err := server.store.MatchAll(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	c.JSON(200, emplsgifts)
}
