package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	DAO *dao.UserDAO
}

func NewUserHandler(dao *dao.UserDAO) *UserHandler {
	return &UserHandler{DAO: dao}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//encrypt the password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 11)
	newUser.Password = string(bytes)

	if err := h.DAO.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Created user": newUser})
}

func (h *UserHandler) ValidateUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error binding": err.Error()})
		return
	}
	db_user, err := h.DAO.GetUser(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	match := bcrypt.CompareHashAndPassword([]byte(db_user.Password), []byte(user.Password))
	if match != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": match.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": "User authenticated succesfully"})
}
