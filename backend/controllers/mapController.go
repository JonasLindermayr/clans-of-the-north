package controllers

import (
	"net/http"
	"strconv"

	"github.com/JonasLindermayr/clans-of-the-north/backend/registry"
	"github.com/gin-gonic/gin"
)

// GetMapController gibt den Kartenausschnitt basierend auf den Center-Koordinaten zurück.
func GetMapController(c *gin.Context) {
	// Hole die Query-Parameter centerX und centerY
	centerX, err := strconv.Atoi(c.DefaultQuery("centerX", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid centerX parameter"})
		return
	}
	centerY, err := strconv.Atoi(c.DefaultQuery("centerY", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid centerY parameter"})
		return
	}

	// Definiere den Bereich der Karte, der zurückgegeben werden soll (z.B. -5 bis 5 für 10x10 Tiles)
	rangeSize := 5 // Hier wird der Range von 5 für ein 10x10-Ausschnitt verwendet
	tilesInRange := make(map[string]interface{})

	// Iteriere durch den Bereich und hole die Tiles
	for x := centerX - rangeSize; x <= centerX + rangeSize; x++ {
		for y := centerY - rangeSize; y <= centerY + rangeSize; y++ {
			key := strconv.Itoa(x) + ":" + strconv.Itoa(y)
			if tile, exists := registry.AllTilesInCache[key]; exists {
				tilesInRange[key] = tile
			}
		}
	}

	// Gib die Tiles als JSON zurück
	c.JSON(http.StatusOK, tilesInRange)
}
