package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/handler"
)

func MapRequisiteRoutes(db *sql.DB, group *gin.RouterGroup) {
	requisiteDAO := dao.NewRequisiteDAO(db)
	requisiteHandler := handler.NewRequisiteHandler(requisiteDAO)

	requisites := group.Group("/requisite")
	{
		requisites.GET("", requisiteHandler.GetRequisites)
		requisites.GET("/:classid/:reqid", requisiteHandler.GetRequisiteByID)
		requisites.POST("", requisiteHandler.CreateRequisite)
		requisites.PUT("/:classid/:reqid", requisiteHandler.UpdateRequisite)
		requisites.DELETE("/:classid/:reqid", requisiteHandler.DeleteRequisite)
	}
}
