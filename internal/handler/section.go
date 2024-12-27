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

type SectionHandler struct {
	DAO *dao.SectionDAO
}

func NewSectionHandler(dao *dao.SectionDAO) *SectionHandler {
	return &SectionHandler{DAO: dao}
}

func (h *SectionHandler) GetSections(c *gin.Context) {
	sections, err := h.DAO.GetSections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sections)
}

func (h *SectionHandler) GetSectionByID(c *gin.Context) {

	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric section id given"})
		return
	}

	section, err := h.DAO.GetSectionByID(idParam)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no section found with ID: %d", idParam)})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to execute query: %s", err)})
		return
	}

	c.JSON(http.StatusOK, section)
}

func (h *SectionHandler) CreateSection(c *gin.Context) {
	var newSection model.Section

	if err := c.Bind(&newSection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error with request body: %s", err)})
		return
	}

	if err := h.DAO.CreateSection(&newSection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Created section": newSection})
}

func (h *SectionHandler) UpdateSection(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non numeric section id given"})
		return
	}

	var updatedSection model.Section

	if err := c.Bind(&updatedSection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.DAO.UpdateSection(&updatedSection, idParam); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no section found with ID %d to update", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated Section": updatedSection})
}

func (h *SectionHandler) DeleteSection(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non numeric section id given"})
		return
	}

	deletedSection, err := h.DAO.DeleteSection(idParam)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no section found with ID %d to delete", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"deleted section": deletedSection})
}
