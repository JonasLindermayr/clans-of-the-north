package functions

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func FindNextVillageSpot(isForsaken bool) *models.Tile {

	dx, dy := 0, -1
	x, y := 0, 0
	max := utils.Config.WorldSize * 2

	for i := 0; i < max*max; i++ {

		tile := registry.GetTile(x, y)

		if tile != nil && tile.Terrain == "village_spot" && tile.VillageID == nil {

			if isForsaken {

				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						neighborTile := registry.GetTile(x+dx, y+dy)
						if neighborTile != nil && neighborTile.Terrain == "plains" {
							return tile
						}
					}
				}
			} else {
				return tile
			}
		}

		if x == y || (x < 0 && x == -y) || (x > 0 && x == 1-y) {
			dx, dy = -dy, dx
		}
		x += dx
		y += dy
	}

	return nil
}
