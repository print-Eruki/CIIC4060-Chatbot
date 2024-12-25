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

type MeetingHandler struct {
	DAO *dao.MeetingDAO
}

func NewMeetingHandler(dao *dao.MeetingDAO) *MeetingHandler {
	return &MeetingHandler{DAO: dao}
}

func (h *MeetingHandler) GetMeetings(c *gin.Context) {
	meetings, err := h.DAO.GetMeetings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, meetings)
}

func (h *MeetingHandler) GetMeetingByID(c *gin.Context) {

	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric meeting id given"})
		return
	}

	meeting, err := h.DAO.GetMeetingByID(idParam)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no meeting found with ID: %d", idParam)})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to execute query: %s", err)})
		return
	}

	c.JSON(http.StatusOK, meeting)
}

func (h *MeetingHandler) CreateMeeting(c *gin.Context) {
	var newMeeting model.Meeting

	if err := c.Bind(&newMeeting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error with request body: %s", err)})
		return
	}

	if err := h.DAO.CreateMeeting(&newMeeting); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Created meeting": newMeeting})
}

func (h *MeetingHandler) UpdateMeeting(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non numeric meeting id given"})
		return
	}

	var updatedMeeting model.Meeting

	if err := c.Bind(&updatedMeeting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := h.DAO.UpdateMeeting(&updatedMeeting, idParam); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no meeting found with ID %d to update", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated Meeting": updatedMeeting})
}

func (h *MeetingHandler) DeleteMeeting(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non numeric meeting id given"})
		return
	}

	deletedMeeting, err := h.DAO.DeleteMeeting(idParam)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no meeting found with ID %d to delete", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"deleted meeting": deletedMeeting})
}
