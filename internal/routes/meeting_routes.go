package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/handler"
)

func MapMeetingRoutes(db *sql.DB, group *gin.RouterGroup) {
	meetingDAO := dao.NewMeetingDAO(db)
	meetingHandler := handler.NewMeetingHandler(meetingDAO)

	meetings := group.Group("/meeting")
	{
		meetings.GET("", meetingHandler.GetMeetings)
		meetings.GET("/:id", meetingHandler.GetMeetingByID)
		meetings.POST("", meetingHandler.CreateMeeting)
		meetings.PUT("/:id", meetingHandler.UpdateMeeting)
		meetings.DELETE("/:id", meetingHandler.DeleteMeeting)
	}
}
