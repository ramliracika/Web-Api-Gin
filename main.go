package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ramliracika/Web-Api-Go/handler"
	"github.com/ramliracika/Web-Api-Go/laptop"
	"github.com/ramliracika/Web-Api-Go/repository"
	"github.com/ramliracika/Web-Api-Go/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/web_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection to Database failed")

	}
	fmt.Println("Connection to Database Succeed")
	db.AutoMigrate(&laptop.Gaming{})
	//auto migrate struct to table database

	laptopRepository := repository.NewRepository(db)
	laptopService := service.NewService(laptopRepository)
	laptopHandler := handler.NewHandler(laptopService)

	route := gin.Default() //delcare route gin

	v1 := route.Group("/v1") //versionin Api

	v1.GET("/", laptopHandler.MainHandler)
	v1.GET("/hello", laptopHandler.HelloHandler)
	v1.GET("/laptop/:id/:brand", laptopHandler.ParamsHandler)
	v1.GET("/stock", laptopHandler.QueryHandler)
	v1.POST("/laptop", laptopHandler.PostHandler)
	v1.GET("/laptop", laptopHandler.GetAllHandler)
	v1.GET("/laptop/:id", laptopHandler.GetByIdHandler)
	v1.GET("/laptop/query", laptopHandler.GetByIdQueryHandler)
	v1.PUT("/laptop/:id", laptopHandler.UpdateHandler)
	v1.DELETE("/laptop/:id", laptopHandler.DeleteHandler)

	route.Run()
}
