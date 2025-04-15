package workers

import (
	"fmt"
	"log"
	"time"

	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
	"gorm.io/gorm"
)

func SaveCacheToDB() {

	for {
		time.Sleep(time.Second * 20)

		fmt.Println("Saving cache to DB")

		for i := range registry.AllVillagesInCache {
			village := &registry.AllVillagesInCache[i]

			err := utils.DB.
				Session(&gorm.Session{FullSaveAssociations: true}).
				Save(village).Error

			if err != nil {
				log.Printf("error while saving village %d: %v", village.ID, err)
			}
		}

	}
}