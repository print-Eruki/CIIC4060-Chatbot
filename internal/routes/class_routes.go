package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/handler"
)

func MapClassRoutes(db *sql.DB, group *gin.RouterGroup) {
	classDAO := dao.NewClassDAO(db)
	classHandler := handler.NewClassHandler(classDAO)

	classes := group.Group("/class")
	{
		classes.GET("", classHandler.GetClasses)
		classes.GET("/:id", classHandler.GetClassByID)
		classes.POST("", classHandler.CreateClass)
		classes.PUT("/:id", classHandler.UpdateClass)
		classes.DELETE("/:id", classHandler.DeleteClass)
	}
}
