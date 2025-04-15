package cache

import "github.com/JonasLindermayr/clans-of-the-north/backend/registry"

func hasEnoughResources(villageID uint, cost map[string]int) bool{

	if registry.AllVillagesInCache[villageID].Resources.Wood < cost["wood"] {
		return false
	}
	if registry.AllVillagesInCache[villageID].Resources.Stone < cost["stone"] {
		return false
	}
	if registry.AllVillagesInCache[villageID].Resources.Clay < cost["clay"] {
		return false
	}
	if registry.AllVillagesInCache[villageID].Resources.Grain < cost["grain"] {
		return false
	}			
	if registry.AllVillagesInCache[villageID].Resources.Gold < cost["gold"] {
		return false
	}										

	return true
}	