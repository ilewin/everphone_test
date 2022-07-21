package api

import (
	"github.com/gin-gonic/gin"
	"github.com/transparentideas/everphone_test/db"
	"github.com/transparentideas/everphone_test/middleware"
)

// Server serves HTTP requests for out gifty service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/employees", server.ListEmployees)
	router.GET("/employees/:name", server.GetEmployeeByName)
	router.PATCH("/employee/", server.UpdateEmployee)
	router.POST("/employee/", middleware.AuthUser(), server.AddEmployee)

	router.GET("/gifts", server.ListGifts)
	router.GET("/gifts/:id", server.GetGift)
	router.POST("/gift/", server.AddGift)
	router.PATCH("/gift/", server.UpdateGift)

	router.GET("/match/:name", server.MatchGiftToEmployee)
	router.GET("/match/", server.MatchAll)

	router.GET("/token", server.GetToken)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(e error) gin.H {
	return gin.H{
		"error": e.Error(),
	}
}
