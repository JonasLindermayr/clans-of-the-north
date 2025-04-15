package workers

import (
	"fmt"
	"time"

	"github.com/JonasLindermayr/clans-of-the-north/backend/functions/cache"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	. "github.com/JonasLindermayr/clans-of-the-north/backend/utils/constants"
)

func ResourceProdWorker() {
	ticker := time.NewTicker(10 * time.Second) 
	defer ticker.Stop()

	for tick := range ticker.C {
		fmt.Println("Running resource production tick at", tick)

		for _, village := range registry.AllVillagesInCache {
			now := time.Now()
			delta := now.Sub(village.Resources.LastResourceUpdate).Seconds()

			if delta < 1 {
				continue
			}

			woodcutter := cache.GetBuilding(village, BUILDING_WOOD_CUTTER)
			stone := cache.GetBuilding(village, BUILDING_STONE_QUARRY)
			clay := cache.GetBuilding(village, BUILDING_CLAY_PIT)
			grain := cache.GetBuilding(village, BUILDING_GRAIN_FIELD)
			gold := cache.GetBuilding(village, BUILDING_GOLD_MINE)

			woodProductionPerSecond := (float64(utils.BuildingsConfig.Buildings[BUILDING_WOOD_CUTTER].ProductionRate) * float64(woodcutter.CurrentLevel) * utils.Config.ResourceProductionMultiply) / 60
			stoneProductionPerSecond := (float64(utils.BuildingsConfig.Buildings[BUILDING_STONE_QUARRY].ProductionRate) * float64(stone.CurrentLevel) * utils.Config.ResourceProductionMultiply) / 60
			clayProductionPerSecond := (float64(utils.BuildingsConfig.Buildings[BUILDING_CLAY_PIT].ProductionRate) * float64(clay.CurrentLevel) * utils.Config.ResourceProductionMultiply) / 60
			grainProductionPerSecond := (float64(utils.BuildingsConfig.Buildings[BUILDING_GRAIN_FIELD].ProductionRate) * float64(grain.CurrentLevel) * utils.Config.ResourceProductionMultiply) / 60
			goldProductionPerSecond := (float64(utils.BuildingsConfig.Buildings[BUILDING_GOLD_MINE].ProductionRate) * float64(gold.CurrentLevel) * utils.Config.ResourceProductionMultiply) / 86400

			woodProduced := woodProductionPerSecond * delta
			stoneProduced := stoneProductionPerSecond * delta
			clayProduced := clayProductionPerSecond * delta
			grainProduced := grainProductionPerSecond * delta
			goldProduced := goldProductionPerSecond * delta

			woodProduced += village.Resources.CarryOverWood
			stoneProduced += village.Resources.CarryOverStone
			clayProduced += village.Resources.CarryOverClay
			grainProduced += village.Resources.CarryOverGrain
			goldProduced += village.Resources.CarryOverGold


			if (village.Resources.Wood + int(woodProduced)) >= village.Resources.Storage {
				village.Resources.Wood = village.Resources.Storage
				village.Resources.CarryOverWood = 0
			} else {
				village.Resources.Wood += int(woodProduced)
				village.Resources.CarryOverWood = woodProduced - float64(int(woodProduced))
			}

			if (village.Resources.Stone + int(stoneProduced)) >= village.Resources.Storage {
				village.Resources.Stone = village.Resources.Storage
				village.Resources.CarryOverStone = 0
			} else {
				village.Resources.Stone += int(stoneProduced)
				village.Resources.CarryOverStone = stoneProduced - float64(int(stoneProduced))
			}

			if (village.Resources.Clay + int(clayProduced)) >= village.Resources.Storage {
				village.Resources.Clay = village.Resources.Storage
				village.Resources.CarryOverClay = 0
			} else {
				village.Resources.Clay += int(clayProduced)
				village.Resources.CarryOverClay = clayProduced - float64(int(clayProduced))
			}

			if (village.Resources.Grain + int(grainProduced)) >= village.Resources.Storage {
				village.Resources.Grain = village.Resources.Storage
				village.Resources.CarryOverGrain = 0
			} else {
				village.Resources.Grain += int(grainProduced)	
				village.Resources.CarryOverGrain = grainProduced - float64(int(grainProduced))
			}

			if (village.Resources.Gold + int(goldProduced)) >= village.Resources.Storage {
				village.Resources.Gold = village.Resources.Storage
				village.Resources.CarryOverGold = 0
			} else {
				village.Resources.Gold += int(goldProduced)
				village.Resources.CarryOverGold = goldProduced - float64(int(goldProduced))
			}

			village.Resources.LastResourceUpdate = now
		}
	}
}
