package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/print-Eruki/CIIC4060-chatbot/config"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/routes"
)

func main() {
	db := config.SetupConnection()
	router := gin.Default()
	//cleans incoming request urls, /entity/ -> /entity
	router.RemoveExtraSlash = true

	// setup the routes for every entity
	//use groups instead
	api := router.Group("/datastic_4/")
	{
		routes.MapClassRoutes(db, api)
		routes.MapMeetingRoutes(db, api)
		routes.MapRoomRoutes(db, api)
		routes.MapSectionRoutes(db, api)
		routes.MapUserRoutes(db, api)
		routes.MapRequisiteRoutes(db, api)
	}

	router.Run("localhost:8080")
}
