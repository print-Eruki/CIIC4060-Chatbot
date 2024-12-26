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

type RoomHandler struct {
	DAO *dao.RoomDAO
}

func NewRoomHandler(dao *dao.RoomDAO) *RoomHandler {
	return &RoomHandler{DAO: dao}
}


func (h *RoomHandler) GetRooms(c *gin.Context) {
	rooms, err := h.DAO.GetRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (h *RoomHandler) GetRoomByID(c *gin.Context) {

	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric room id given"})
		return
	}

	room, err := h.DAO.GetRoomByID(idParam)
	if err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no room found with ID: %d", idParam)})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to execute query: %s", err)})
		return
	}
	c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var newRoom model.Room

	if err := c.Bind(&newRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error with request body: %s", err)})
		return
	}

	if err := h.DAO.CreateRoom(&newRoom); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"Created room": newRoom})
}

func (h *RoomHandler) UpdateRoom(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric room id given"})
		return
	}

	var updatedRoom model.Room

	if err := c.Bind(&updatedRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = h.DAO.UpdateRoom(&updatedRoom, idParam); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no room found with ID %d to update", idParam)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Updated room": updatedRoom})
}

func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	idParam, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "non-numeric room id given"})
		return
	}

	deletedRoom, err := h.DAO.DeleteRoom(idParam)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no room found with ID %d to delete", idParam)})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Deleted room": deletedRoom})
}

