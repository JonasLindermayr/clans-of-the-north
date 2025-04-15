package cache

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func UpgradeBuilding(villageID uint, buildingId uint) {

	building := &registry.AllVillagesInCache[villageID].Buildings[buildingId]

	if(building.CurrentLevel < utils.BuildingsConfig.Buildings[building.BuildingID].MaxLevel) {
		building.CurrentLevel++
		if(building.BuildingID == 8 || building.BuildingID == 11) {
			UpdateVillageStorageAndPopulation(villageID)
		}

		UpdateVillagePoints(villageID)
		registry.SaveCacheToDB()
	}

}
