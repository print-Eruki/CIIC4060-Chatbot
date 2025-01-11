package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/handler"
)

func MapRoomRoutes(db *sql.DB, group *gin.RouterGroup) {
	roomDAO := dao.NewRoomDAO(db)
	roomHandler := handler.NewRoomHandler(roomDAO)

	rooms := group.Group("/room")
	{
		rooms.GET("", roomHandler.GetRooms)
		rooms.GET("/:id", roomHandler.GetRoomByID)
		rooms.POST("", roomHandler.CreateRoom)
		rooms.PUT("/:id", roomHandler.UpdateRoom)
		rooms.DELETE("/:id", roomHandler.DeleteRoom)
	}
}
