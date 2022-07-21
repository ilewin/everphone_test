package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/transparentideas/everphone_test/db"
)

func (server *Server) ListEmployees(c *gin.Context) {
	employees, err := server.store.ListEmployees(c.Request.Context(), db.ListEmployeesParams{})
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, employees)
}

func (server *Server) GetEmployeeByName(c *gin.Context) {
	var req db.GetEmployeeByNameParams

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	employees, err := server.store.GetEmployeeByName(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, employees)
}

func (server *Server) UpdateEmployee(c *gin.Context) {
	var req db.UpdateEmployeeParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	employee, err := server.store.UpdateEmployee(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, employee)
}

func (server *Server) AddEmployee(c *gin.Context) {
	var req db.AddEmployeeParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	employees, err := server.store.AddEmployee(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}
	c.JSON(200, employees)
}
