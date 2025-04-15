package cache

import (
	"time"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)


func CreateBuildOrder(villageID uint, buildingID uint) {
	timeDecrease := 0.0
	currentLevel := registry.AllVillagesInCache[villageID].Buildings[buildingID].CurrentLevel
	maxLevel := utils.BuildingsConfig.Buildings[buildingID].MaxLevel

	baseConstructionTime := float64(utils.BuildingsConfig.Buildings[buildingID].ConstructionTime)


	if currentLevel == maxLevel {
		return
	}

	if len(registry.AllVillagesInCache[villageID].BuildQueue) <= 2 {
		return
	}

	if registry.AllVillagesInCache[villageID].BuildQueue[0].BuildingID == buildingID {
		currentLevel++;
	}

	if registry.AllVillagesInCache[villageID].Buildings[buildingID].CurrentLevel >= 1 {
		baseConstructionTime = baseConstructionTime * utils.Config.BuildingConstructionTimeMultiply * 
								float64(currentLevel)
	}


	if registry.AllVillagesInCache[villageID].Buildings[10].CurrentLevel > 1 {
		timeDecrease = float64(registry.AllVillagesInCache[villageID].Buildings[10].CurrentLevel) *
						utils.BuildingsConfig.Buildings[10].BuildingTimeDecrease
	}


	timeToFinish := time.Duration(baseConstructionTime * (1 - timeDecrease)) * time.Second

	buildOrder := &models.BuildOrder{
		VillageID: villageID,
		BuildingID: buildingID,
		StartTime: time.Now(),
		EndTime: time.Now().Add(timeToFinish),
	}

	registry.AllVillagesInCache[villageID].BuildQueue = append(registry.AllVillagesInCache[villageID].BuildQueue, *buildOrder)

}
