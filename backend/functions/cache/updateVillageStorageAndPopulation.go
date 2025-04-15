package cache

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	. "github.com/JonasLindermayr/clans-of-the-north/backend/utils/constants"
)

func UpdateVillageStorageAndPopulation(villageID uint) {

	var storageFloat float64
	var populationFloat float64

	if registry.AllVillagesInCache[villageID].Buildings[BUILDING_STOREHOUSE].CurrentLevel > 1 {
		storageFloat = float64((utils.BuildingsConfig.Buildings[BUILDING_STOREHOUSE].BaseStorage * 
			registry.AllVillagesInCache[villageID].Buildings[BUILDING_STOREHOUSE].CurrentLevel)) *
			(utils.Config.WarehouseStorageMultiply)
	} else {
		storageFloat = float64((utils.BuildingsConfig.Buildings[BUILDING_STOREHOUSE].BaseStorage *
			registry.AllVillagesInCache[villageID].Buildings[BUILDING_STOREHOUSE].CurrentLevel))
	}

	if registry.AllVillagesInCache[villageID].Buildings[BUILDING_HALL_OF_PEOPLE].CurrentLevel > 1 {
		populationFloat = float64((utils.BuildingsConfig.Buildings[BUILDING_HALL_OF_PEOPLE].PopulationBaseValue * 
			registry.AllVillagesInCache[villageID].Buildings[BUILDING_HALL_OF_PEOPLE].CurrentLevel)) *
			(utils.Config.PopulationCapacityIncrease)
	} else {
		populationFloat = float64((utils.BuildingsConfig.Buildings[BUILDING_HALL_OF_PEOPLE].PopulationBaseValue * 
			registry.AllVillagesInCache[villageID].Buildings[BUILDING_HALL_OF_PEOPLE].CurrentLevel))
	}

	registry.AllVillagesInCache[villageID].Resources.Storage = int(storageFloat)
	registry.AllVillagesInCache[villageID].Resources.Population = int(populationFloat)
	
}