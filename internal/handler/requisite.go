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

type RequisiteHandler struct {
	DAO *dao.RequisiteDAO
}

func NewRequisiteHandler(dao *dao.RequisiteDAO) *RequisiteHandler {
	return &RequisiteHandler{DAO: dao}
}

func (h *RequisiteHandler) GetRequisites(c *gin.Context) {
	requisites, err := h.DAO.GetRequisites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, requisites)
}

func (h *RequisiteHandler) GetRequisiteByID(c *gin.Context) {

	cIdParam, err := strconv.ParseUint(c.Param("classid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric requisite id given"})
		return
	}
	rIdParam, err := strconv.ParseUint(c.Param("reqid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric reqid given"})
		return
	}

	requisite, err := h.DAO.GetRequisiteByID(cIdParam, rIdParam)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no requisite found with classid %d and reqid %d", cIdParam, rIdParam)})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to execute query: %s", err)})
		return
	}

	c.JSON(http.StatusOK, requisite)
}

func (h *RequisiteHandler) CreateRequisite(c *gin.Context) {
	var newRequisite model.Requisite

	if err := c.Bind(&newRequisite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error with request body: %s", err)})
		return
	}

	if err := h.DAO.CreateRequisite(&newRequisite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Created requisite": newRequisite})
}

func (h *RequisiteHandler) UpdateRequisite(c *gin.Context) {
	cIdParam, err := strconv.ParseUint(c.Param("classid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric requisite id given"})
		return
	}
	rIdParam, err := strconv.ParseUint(c.Param("reqid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric reqid given"})
		return
	}
	var updatedRequisite model.Requisite

	if err := c.Bind(&updatedRequisite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.DAO.UpdateRequisite(&updatedRequisite, cIdParam, rIdParam); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no requisite found with classid %d and reqid %d to update", cIdParam, rIdParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated Requisite": updatedRequisite})
}

func (h *RequisiteHandler) DeleteRequisite(c *gin.Context) {
	cIdParam, err := strconv.ParseUint(c.Param("classid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric requisite id given"})
		return
	}
	rIdParam, err := strconv.ParseUint(c.Param("reqid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric reqid given"})
		return
	}

	deletedRequisite, err := h.DAO.DeleteRequisite(cIdParam, rIdParam)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no requisite found with classid %d and reqid %d to delete", cIdParam, rIdParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"deleted requisite": deletedRequisite})
}
