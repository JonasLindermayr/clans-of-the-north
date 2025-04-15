package utils

import (
	"encoding/json"
	"log"
	"os"
)

type BuildingInfo struct {
    ID                   int               `json:"id"`
    Name                 string            `json:"name"`
    Description          string            `json:"description"`
    ResourceType         string            `json:"resourceType,omitempty"`
    ProductionRate       int                `json:"productionRate,omitempty"`
    MaxLevel             int               `json:"maxLevel"`
    StartLevel           int               `json:"startLevel"`
    ConstructionTime     int               `json:"constructionTime"`
    Cost                 map[string]int    `json:"upgradeCost"`
    PopulationBaseValue  int               `json:"populationBaseValue,omitempty"`
    DefenseBonus         float64               `json:"defenseBonus,omitempty"`
    BuildingTimeDecrease float64               `json:"buildingTimeDecrease,omitempty"`
    Points              int               `json:"points"`
    BaseStorage         int               `json:"baseStorage,omitempty"`
}


type BuildingsInfoConfig struct {
    Buildings []BuildingInfo `json:"buildings"`
}

var BuildingsConfig *BuildingsInfoConfig

func LoadBuildingInfos() {
	data, err := os.ReadFile("./buildings.json")
	if err != nil {
		log.Fatal("could not read buildings file")
	}

	if err := json.Unmarshal(data, &BuildingsConfig); err != nil {
		log.Fatal("could not parse buildings file")
	}
}
