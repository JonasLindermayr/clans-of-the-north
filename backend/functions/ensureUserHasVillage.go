package functions

import (
	"github.com/JonasLindermayr/clans-of-the-north/backend/models"
	"github.com/JonasLindermayr/clans-of-the-north/backend/utils"
)

func EnsureUserHasVillage(uuid uint, username string) {
	var village models.Village

	err := utils.DB.Where("user_id =?", uuid).First(&village).Error
	if err == nil {
		return
	}

	createInitialVillage(uuid, username + "'s Village")

}