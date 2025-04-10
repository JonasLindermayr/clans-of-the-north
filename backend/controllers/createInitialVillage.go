package controllers

import (
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func createInitialPlayerVillage(uuid uint, username string) {

	village := models.Village{
		Name: username + "'s Village",
		UserID: uuid,
		CordX: 0,
		CordY: 0,
		Points: 10,
	}

	
	if err := utils.DB.Create(&village).Error; err != nil {
		log.Fatal("error while creating village")
		return
	}
	
	createInitialResources(village.ID)

}