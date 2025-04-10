package controllers

import (
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func createInitialResources(villageID uint) {
	resources := models.Resources {
		VillageID: villageID,
		Wood: 100,
		Stone: 100,
		Clay: 100,
		Grain: 100,
	}

	
	if err := utils.DB.Create(&resources).Error; err != nil {
		log.Fatal("error while creating resources for village")
		return
	}
}