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

	classDAO := dao.NewClassDAO(db)
	classHandler := handler.NewClassHandler(classDAO)
	
	router.GET("/datastic_4/class", classHandler.GetClasses)

	router.Run("localhost:8080")
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
