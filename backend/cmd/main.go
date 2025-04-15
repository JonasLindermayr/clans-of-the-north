package main

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/controllers"
	"github.com/JonasLindermayr/clans-of-the-north/backend/middleware"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	"github.com/JonasLindermayr/clans-of-the-north/backend/workers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnvs()
	utils.ConnectDB()
	utils.LoadConfig()
	utils.LoadBuildingInfos()
	registry.LoadIntoCache()
}

func main() {

	router := gin.Default()


	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.LoginUser)
	router.GET("/user/profile", middleware.CheckAuth, controllers.GetUserProfile)

	go workers.StartWorkers()

	router.Run(":8080")
}