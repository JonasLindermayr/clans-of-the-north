package functions

import (
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func CreateInitialBuildings(villageID uint) {
	var buildings []models.Buildings

	for _, building := range utils.BuildingsConfig.Buildings {

		if err := utils.DB.Where("village_id = ? AND building_id = ?", villageID, building.ID).First(&models.Buildings{}).Error; err == nil {
			continue
		}

		initBuilding := models.Buildings {
			VillageID: villageID,
			BuildingID: building.ID,
			Type: building.Name,
			CurrentLevel: building.StartLevel,
		}

		buildings = append(buildings, initBuilding)
	}


	if buildings == nil {
		return
	}

	if err := utils.DB.Create(&buildings).Error; err != nil {
		log.Fatal(err.Error())
	}

}
