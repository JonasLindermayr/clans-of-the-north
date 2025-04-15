package functions

import "fmt"

func PlaceVillage() (x, y int) {
	centerX, centerY := 0, 0
	maxRadius := 120

	for r := 0; r < maxRadius; r++ {
		for dx := -r; dx <= r; dx++ {
			for dy := -r; dy <= r; dy++ {
				x := centerX + dx
				y := centerY + dy

				fmt.Println(x, y)

				/*
				if IsTileFree(x, y) {
					return x, y
				*/
			}
		}
	}

	return 0, 0 
}

/*
func IsTileFree(x, y int) bool {
	tile, exists := registry.MapTiles[[2]int{x, y}]
	if !exists || tile.VillageID == nil {
		return true
	}
	return false
}
	*/