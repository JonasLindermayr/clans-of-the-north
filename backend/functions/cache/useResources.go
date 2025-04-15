package cache

import "github.com/JonasLindermayr/clans-of-the-north/backend/registry"

func useResources(villageID uint, cost map[string]int) bool{

	if !hasEnoughResources(villageID, cost) {
		return false
	}
	
	registry.AllVillagesInCache[villageID].Resources.Wood -= cost["wood"]
	registry.AllVillagesInCache[villageID].Resources.Stone -= cost["stone"]
	registry.AllVillagesInCache[villageID].Resources.Clay -= cost["clay"]
	registry.AllVillagesInCache[villageID].Resources.Grain -= cost["grain"]
	registry.AllVillagesInCache[villageID].Resources.Gold -= cost["gold"]

	return true
}