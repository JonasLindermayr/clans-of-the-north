package functions

import (
	"fmt"
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/functions/cache"
	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func createInitialVillage(uuid uint, villageName string) {


	tile := FindNextVillageSpot(false)
	if tile == nil {
		log.Fatal("error while finding next village spot")
		return
	}

	village := models.Village{
		Name: villageName,
		UserID: uuid,
		CordX: tile.X,
		CordY: tile.Y,
		Points: 0,
	}


	if err := utils.DB.Create(&village).Error; err != nil {
		log.Fatal("error while creating village")
		return
	}
	
	tile.VillageID = &village.ID
	registry.AllTilesInCache[fmt.Sprintf("%d:%d", tile.X, tile.Y)] = tile

	CreateInitialResources(village.ID)
	CreateInitialBuildings(village.ID)

	registry.ReloadCache()

	cache.UpdateVillagePoints(village.ID)
	cache.UpdateVillageStorageAndPopulation(village.ID)

	registry.SaveCacheToDB()
}
