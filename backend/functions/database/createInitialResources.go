package functions

import (
	"log"
	"time"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func CreateInitialResources(villageID uint) {
	resources := models.Resources {
		VillageID: villageID,
		Wood: 100,
		Stone: 100,
		Clay: 100,
		Grain: 100,
		LastResourceUpdate: time.Now(),
		CarryOverWood: 0,
		CarryOverStone: 0,
		CarryOverClay: 0,
		CarryOverGrain: 0,
		CarryOverGold: 0,
		ClaimedPopulation: 0,
		Population: 0,
		Storage: 0,
	}

	
	if err := utils.DB.Create(&resources).Error; err != nil {
		log.Fatal("error while creating resources for village")
		return
	}
}