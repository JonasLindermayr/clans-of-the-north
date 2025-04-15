package utils

import (
	"log"
	"math/rand"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
)

func GenerateMap(){

	var tiles []models.Tile

	for x := -Config.WorldSize; x <= Config.WorldSize; x++ {
		for y := -Config.WorldSize; y <= Config.WorldSize; y++ {

			terrain := "plains"
			randNum := rand.Intn(100)
	
			if randNum < 10 {
				terrain = "forest"  
			} else if randNum < 20 {
				terrain = "mountain" 
			} else if randNum < 25 {
				terrain = "lake" 
			} else if randNum < 45 {
				terrain = "village_spot"
			} 
			
			tile := models.Tile{
				X: x,
				Y: y,
				Terrain: terrain,
				VillageID: nil,
			}

			tiles = append(tiles, tile)
		}
	}

	const batchSize = 1000 

	if err := DB.CreateInBatches(&tiles, batchSize).Error; err != nil {
		log.Fatalf("Fehler beim Map-Insert: %v", err)
	}

}