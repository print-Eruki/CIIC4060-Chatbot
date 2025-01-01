package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
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

func (h *ClassHandler) GetClassByID(c *gin.Context) {

	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric class id given"})
		return
	}

	class, err := h.DAO.GetClassByID(idParam)
	if err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no class found with ID: %d", idParam)})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to execute query: %s", err)})
		return
	}
	c.JSON(http.StatusOK, class)
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var newClass model.Class

	if err := c.Bind(&newClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error with request body: %s", err)})
		return
	}

	if err := h.DAO.CreateClass(&newClass); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"Created class": newClass})
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric class id given"})
		return
	}

	var updatedClass model.Class

	if err := c.Bind(&updatedClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.DAO.UpdateClass(&updatedClass, idParam); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no class found with ID %d to update", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated class": updatedClass})
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric class id given"})
		return
	}

	deletedClass, err := h.DAO.DeleteClass(idParam)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no class found with ID %d to delete", idParam)})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Deleted class": deletedClass})
}
