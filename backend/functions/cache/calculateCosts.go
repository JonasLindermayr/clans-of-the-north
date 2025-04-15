package cache

func calculateCosts(cost map[string]int, factor float64) map[string]int {
	newCost := make(map[string]int)
	for resource, amount := range cost {
		floatCost := float64(amount) * factor
		newCost[resource] = int(floatCost)
	}
	return newCost
}