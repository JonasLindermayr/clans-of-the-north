package registry

import (
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

var AllVillagesInCache []models.Village

func LoadIntoCache() {

	err := utils.DB.
				Preload("Resources").
				Preload("Buildings").
				Find(&AllVillagesInCache).Error
	if err != nil {
		log.Fatal("error loading Villages in to Cache")
	}

}

func ReloadCache() {
	AllVillagesInCache = nil
	LoadIntoCache()
}