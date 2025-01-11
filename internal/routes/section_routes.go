package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/handler"
)

func MapSectionRoutes(db *sql.DB, group *gin.RouterGroup) {
	sectionDAO := dao.NewSectionDAO(db)
	sectionHandler := handler.NewSectionHandler(sectionDAO)

	sections := group.Group("/section")
	{
		sections.GET("", sectionHandler.GetSections)
		sections.GET("/:id", sectionHandler.GetSectionByID)
		sections.POST("", sectionHandler.CreateSection)
		sections.PUT("/:id", sectionHandler.UpdateSection)
		sections.DELETE("/:id", sectionHandler.DeleteSection)
	}
}
