package cache

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func UpdateVillageStorageAndPopulation(villageID uint) {

	var storageFloat float64
	var populationFloat float64

	if registry.AllVillagesInCache[villageID].Buildings[8].CurrentLevel > 1 {
		storageFloat = float64((utils.BuildingsConfig.Buildings[8].BaseStorage * 
			registry.AllVillagesInCache[villageID].Buildings[8].CurrentLevel)) *
			(utils.Config.WarehouseStorageMultiply)
	} else {
		storageFloat = float64((utils.BuildingsConfig.Buildings[8].BaseStorage *
			registry.AllVillagesInCache[villageID].Buildings[8].CurrentLevel))
	}

	if registry.AllVillagesInCache[villageID].Buildings[11].CurrentLevel > 1 {
		populationFloat = float64((utils.BuildingsConfig.Buildings[11].PopulationBaseValue * 
			registry.AllVillagesInCache[villageID].Buildings[11].CurrentLevel)) *
			(utils.Config.PopulationCapacityIncrease)
	} else {
		populationFloat = float64((utils.BuildingsConfig.Buildings[11].PopulationBaseValue * 
			registry.AllVillagesInCache[villageID].Buildings[11].CurrentLevel))
	}

	registry.AllVillagesInCache[villageID].Resources.Storage = int(storageFloat)
	registry.AllVillagesInCache[villageID].Resources.Population = int(populationFloat)
	
}