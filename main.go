package main

import (
	"log"
	"net/http"
	"yearbook/component/appctx"
	"yearbook/middleware"
	todotransport "yearbook/modules/todo/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	dsn := "sqlserver://sa:123456@localhost:1433?database=todoApp"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := runServer(db); err != nil {
		log.Fatalln(err)
	}
}

func runServer(db *gorm.DB) error {
	appCtx := appctx.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	todos := r.Group("/todos")
	{
		todos.GET("/:id", todotransport.GetTodo(appCtx))
	}
	return r.Run()
}
