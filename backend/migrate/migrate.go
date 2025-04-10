package main

import (
	"fmt"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func init() {
	utils.LoadEnvs();
	utils.ConnectDB();
}

func main() {
	err := utils.DB.AutoMigrate(&models.User{}, &models.Village{}, &models.Resources{})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successful migrated database")
}