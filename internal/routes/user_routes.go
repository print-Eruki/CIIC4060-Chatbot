package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/handler"
)

func MapUserRoutes(db *sql.DB, group *gin.RouterGroup) {
	userDAO := dao.NewUserDAO(db)
	userHandler := handler.NewUserHandler(userDAO)

	users := group.Group("")
	{
		users.POST("/signup", userHandler.CreateUser)
		users.POST("/login", userHandler.ValidateUser)
	}
}
