package registry

import (
	"fmt"
	"log"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	"gorm.io/gorm"
)

var AllVillagesInCache map[uint]*models.Village
var AllTilesInCache map[string]*models.Tile

func LoadIntoCache() {

	AllVillagesInCache = make(map[uint]*models.Village)
	AllTilesInCache = make(map[string]*models.Tile)

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



	var tiles []models.Tile
	err = utils.DB.Find(&tiles).Error
	if err != nil {
		log.Fatal("error loading Tiles in to Cache")
	}

	for i := range tiles {
		tile := &tiles[i]
		key := fmt.Sprintf("%d:%d", tile.X, tile.Y)
		AllTilesInCache[key] = tile
	}

}

func ReloadCache() {
	SaveCacheToDB()
	AllVillagesInCache = nil
	AllTilesInCache = nil
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


	for _, tile := range AllTilesInCache {
		err := utils.DB.
			Session(&gorm.Session{FullSaveAssociations: true}).
			Save(tile).Error

		if err != nil {
			log.Printf("error while saving tile %d:%d: %v", tile.X, tile.Y, err)
		}
	}
}

func GetTile(x, y int) *models.Tile {
	key := fmt.Sprintf("%d:%d", x, y)
	return AllTilesInCache[key]
}