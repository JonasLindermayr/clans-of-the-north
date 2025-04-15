package main

import (
	"fmt"
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func init() {
	utils.LoadEnvs();
	utils.LoadConfig()
	utils.ConnectDB();
}

func main() {
	err := utils.DB.AutoMigrate(&models.User{}, &models.Village{}, &models.Resources{}, &models.Buildings{})

	if err != nil {
		fmt.Println(err)
	}
	
	forsakenUser := models.User{
		ID: utils.Config.UUIDForsakenVillages,
		Username: "Forsakens",
		Email: "system@cotn.com",
		Password: "Ff5OvT6mG/lsvC7rkHe0jt5qALB72oYFTH6Ru0emnlM=",
	}
	
	if err := utils.DB.Create(&forsakenUser).Error; err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Successful migrated database")
}