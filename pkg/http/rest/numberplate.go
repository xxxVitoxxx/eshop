package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
)

// NumberPlateHandler _
type NumberPlateHandler struct {
	s *numberplate.Service
}

// NewNumberPlateHandler will return instance
func NewNumberPlateHandler(s *numberplate.Service) *NumberPlateHandler {
	return &NumberPlateHandler{s}
}

func (h *NumberPlateHandler) Route(r *gin.Engine) {
	api := r.Group("/eshop_api/number_plate")

	api.PUT("/condition/:store_name", h.PutConditionByStoreName)
	api.POST("/condition", h.CreateCondition)
}

// PutConditionByStoreName _
func (h *NumberPlateHandler) PutConditionByStoreName(c *gin.Context) {
	storeName := c.Param("store_name")
	req := numberplate.PutCondition{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.s.PutConditionByStoreName(storeName, req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

// CreateCondition _
func (h *NumberPlateHandler) CreateCondition(c *gin.Context) {
	var req numberplate.Condition
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.s.CreateCondition(req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
