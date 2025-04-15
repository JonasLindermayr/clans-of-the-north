package cache

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func UpdateVillagePoints(villageID uint) {
	
	var points int
	for _, building := range registry.AllVillagesInCache[villageID].Buildings {
		points += utils.BuildingsConfig.Buildings[building.BuildingID].Points * building.CurrentLevel
	}

	registry.AllVillagesInCache[villageID].Points = points
}
