package registry

import (
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	"gorm.io/gorm"
)

var AllVillagesInCache map[uint]*models.Village

func LoadIntoCache() {

	AllVillagesInCache = make(map[uint]*models.Village)

	var villages []models.Village
	err := utils.DB.
				Preload("BuildQueue").
				Preload("Resources").
				Preload("Buildings").
				Find(&villages).Error
	if err != nil {
		log.Fatal("error loading Villages in to Cache")
	}

	for i := range villages {
		v := &villages[i]
		AllVillagesInCache[v.ID] = v
	}


}

func ReloadCache() {
	SaveCacheToDB()
	AllVillagesInCache = nil
	LoadIntoCache()
}

func SaveCacheToDB() {
	for _, village := range AllVillagesInCache {
	
		err := utils.DB.
			Session(&gorm.Session{FullSaveAssociations: true}).
			Save(village).Error

		if err != nil {
			log.Printf("error while saving village %d: %v", village.ID, err)
		}
	}
}