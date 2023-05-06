package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jai-rathore/lists/logger"
	"github.com/jai-rathore/lists/models"
	"github.com/jai-rathore/lists/services"
)

// ListHandler handles list requests
type ListHandler struct {
	listService services.IListService
}

// NewListHandler creates a new list handler
func NewListHandler(listService services.IListService, r *gin.Engine) *ListHandler {
	handler := &ListHandler{listService: listService}

	r.GET("/lists/:id", handler.GetListById)
	r.POST("/lists", handler.CreateList)

	return handler
}

// GetListById godoc
// @Summary Get List By Id
// @Schemes
// @Description get list by id
// @Tags lists
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {object} models.List
// @Router /lists/{id} [get]
func (l *ListHandler) GetListById(c *gin.Context) {
	id := c.Param("id")
	id_num, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("not a valid int id " + err.Error())
		c.JSON(400, err.Error())
		return
	}
	list, err := l.listService.GetListById(id_num)
	if err != nil {
		logger.Error("failed to get list by id " + err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, list)
}

// CreateList godoc
// @Summary Create List
// @Schemes
// @Description create list
// @Tags lists
// @Accept json
// @Produce json
// @Param list body models.List true "List"
// @Success 200 {object} models.List
// @Router /lists [post]
func (l *ListHandler) CreateList(c *gin.Context) {
	var list models.List
	if err := c.ShouldBindJSON(&list); err != nil {
		logger.Error("failed to bind json " + err.Error())
		c.JSON(400, err.Error())
		return
	}
	listResp, err := l.listService.CreateList(&list)
	if err != nil {
		logger.Error("failed to create list " + err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, listResp)
}
