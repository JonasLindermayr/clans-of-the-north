package cache

import (
	"math"
	"time"

	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	. "github.com/JonasLindermayr/clans-of-the-north/backend/utils/constants"
)


func CreateBuildOrder(villageID uint, buildingID uint) {
	timeDecrease := 0.0

	building := GetBuilding(registry.AllVillagesInCache[villageID], int(buildingID))
	buildingLonghouse := GetBuilding(registry.AllVillagesInCache[villageID], BUILDING_LONGHOUSE)

	if building == nil {
		return
	}
	currentLevel := building.CurrentLevel
	maxLevel := utils.BuildingsConfig.Buildings[building.BuildingID].MaxLevel

	baseConstructionTime := float64(utils.BuildingsConfig.Buildings[building.BuildingID].ConstructionTime)

	if len(registry.AllVillagesInCache[villageID].BuildQueue) >= utils.Config.MaxBuildOrders {
		return
	}

	if currentLevel == maxLevel {
		return
	}

	if len(registry.AllVillagesInCache[villageID].BuildQueue) > 0 && registry.AllVillagesInCache[villageID].BuildQueue[0].BuildingID == buildingID {
		currentLevel++
	}
	
	costs := calculateCosts(utils.BuildingsConfig.Buildings[building.BuildingID].Cost,
		math.Pow(utils.Config.BuildingConstructionCostMultiply, float64(currentLevel)),
	   )
	   
	if !useResources(villageID, costs) {
	return
}
	if building.CurrentLevel >= 1 {
		baseConstructionTime = baseConstructionTime * utils.Config.BuildingConstructionTimeMultiply * 
								float64(currentLevel)
	}


	if buildingLonghouse.CurrentLevel > 1 {
		timeDecrease = float64(buildingLonghouse.CurrentLevel) *
						utils.BuildingsConfig.Buildings[BUILDING_LONGHOUSE].BuildingTimeDecrease
	}


	timeToFinish := time.Duration(baseConstructionTime * (1 - timeDecrease)) * time.Second

	buildOrder := &models.BuildOrder{
		VillageID: villageID,
		BuildingID: buildingID,
		StartTime: time.Now(),
		EndTime: time.Now().Add(timeToFinish),
	}

	registry.AllVillagesInCache[villageID].BuildQueue = append(registry.AllVillagesInCache[villageID].BuildQueue, *buildOrder)
	registry.SaveCacheToDB()
}
