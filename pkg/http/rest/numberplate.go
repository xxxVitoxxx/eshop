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

	api.PUT("/condition", h.PutCondition)
}

// PutCondition _
func (h *NumberPlateHandler) PutCondition(c *gin.Context) {
	req := numberplate.PutCondition{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.s.PutCondition(req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
