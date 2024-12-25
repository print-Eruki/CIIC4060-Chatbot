package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	config "github.com/print-Eruki/CIIC4060-chatbot/configs"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/dao"
	"github.com/print-Eruki/CIIC4060-chatbot/internal/handler"
)

func main() {
	db := setupConnection()
	router := gin.Default()
	//cleans incoming request urls, /entity/ -> /entity
	router.RemoveExtraSlash = true

	mapClassRoutes(db, router)
	mapMeetingRoutes(db, router)

	router.Run("localhost:8080")
}

// Maps the CRUD classes to their respective routes
func mapClassRoutes(db *sql.DB, router *gin.Engine) {
	classDAO := dao.NewClassDAO(db)
	classHandler := handler.NewClassHandler(classDAO)

	router.GET("/datastic_4/class", classHandler.GetClasses)
	router.GET("/datastic_4/class/:id", classHandler.GetClassByID)
	router.POST("/datastic_4/class", classHandler.CreateClass)
	router.PUT("/datastic_4/class/:id", classHandler.UpdateClass)
	router.DELETE("datastic_4/class/:id", classHandler.DeleteClass)
}

func mapMeetingRoutes(db *sql.DB, router *gin.Engine) {
	meetingDAO := dao.NewMeetingDAO(db)
	meetingHandler := handler.NewMeetingHandler(meetingDAO)

	router.GET("/datastic_4/meeting", meetingHandler.GetMeetings)
	router.GET("/datastic_4/meeting/:id", meetingHandler.GetMeetingByID)
	router.POST("/datastic_4/meeting", meetingHandler.CreateMeeting)
	router.PUT("/datastic_4/meeting/:id", meetingHandler.UpdateMeeting)
	router.DELETE("datastic_4/meeting/:id", meetingHandler.DeleteMeeting)
}

// Establishes a connection to a postgres database
//
// Returns a sql.DB
func setupConnection() *sql.DB {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Db config didn't load, Error: %s", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Database connection failed, details: %s", err)
	}

	return db
}
