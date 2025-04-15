package cache

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	. "github.com/JonasLindermayr/clans-of-the-north/backend/utils/constants"
)

func UpgradeBuilding(villageID uint, buildingID uint) {

	building := GetBuilding(registry.AllVillagesInCache[villageID], int(buildingID))

	if(building.CurrentLevel < utils.BuildingsConfig.Buildings[building.BuildingID].MaxLevel) {
		building.CurrentLevel++
		if(building.BuildingID == BUILDING_STOREHOUSE || building.BuildingID == BUILDING_HALL_OF_PEOPLE) {
			UpdateVillageStorageAndPopulation(villageID)
		}

		UpdateVillagePoints(villageID)
		registry.SaveCacheToDB()
	}

}
