package functions

import (
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/functions/cache"
	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func createInitialVillage(uuid uint, villageName string) {

	village := models.Village{
		Name: villageName,
		UserID: uuid,
		CordX: 0,
		CordY: 0,
		Points: 0,
	}


	if err := utils.DB.Create(&village).Error; err != nil {
		log.Fatal("error while creating village")
		return
	}
	

	CreateInitialResources(village.ID)
	CreateInitialBuildings(village.ID)

	registry.ReloadCache()

	cache.UpdateVillagePoints(village.ID)
	cache.UpdateVillageStorageAndPopulation(village.ID)

	registry.SaveCacheToDB()
}
