package workers

import (
	"time"

	"github.com/JonasLindermayr/clans-of-the-north/backend/functions/cache"
	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func BuildingQueueWorker() {
	ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()

	for range ticker.C {
		for _, village := range registry.AllVillagesInCache {

			if len(village.BuildQueue) == 0 {
				continue
			} 

			current := village.BuildQueue[0]

			if time.Now().After(current.EndTime) {
				cache.UpgradeBuilding(current.VillageID, current.BuildingID)
				utils.DB.Where("id = ? AND village_id = ? AND building_id = ? AND start_time = ? AND end_time = ?", 
							current.ID, current.VillageID, current.BuildingID, current.StartTime, current.EndTime).Delete(&models.BuildOrder{})

				if len(village.BuildQueue) > 0 {
					village.BuildQueue = village.BuildQueue[1:]
				}
			}

		}
	}
}