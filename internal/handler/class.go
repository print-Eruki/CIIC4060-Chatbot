package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
)

type ClassHandler struct {
	DAO *dao.ClassDAO
}

func NewClassHandler(dao *dao.ClassDAO) *ClassHandler {
	return &ClassHandler{DAO: dao}
}

func (h *ClassHandler) GetClasses(c *gin.Context) {
	classes, err := h.DAO.GetClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}
