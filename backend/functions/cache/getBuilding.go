package cache

import "github.com/JonasLindermayr/clans-of-the-north/backend/models"

func GetBuilding(village *models.Village, buildingID int) *models.Buildings{
	for i := range village.Buildings {
		if village.Buildings[i].BuildingID == buildingID {
			return &village.Buildings[i]
		}
	}
	return nil
}