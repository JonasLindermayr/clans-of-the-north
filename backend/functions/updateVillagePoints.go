package functions

import (
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func UpdateVillagePoints(villageID uint) {
	var village models.Village
	var buildings []models.Buildings

	if err := utils.DB.First(&village, villageID).Error; err != nil {
		log.Println("Village not found:", err)
		return
	}

	if err := utils.DB.Where("village_id = ?", villageID).Find(&buildings).Error; err != nil {
		log.Println("Error fetching buildings:", err)
		return
	}

	var points int
	for _, building := range buildings {
		points += utils.BuildingsConfig.Buildings[building.BuildingID].Points * building.CurrentLevel
	}

	village.Points = points
	utils.DB.Save(&village)
}
