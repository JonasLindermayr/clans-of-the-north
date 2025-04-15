package utils

import (
	"encoding/json"
	"fmt"
	"os"
)
type GameConfig struct {
    UUIDForsakenVillages            uint  `json:"UUID_FORSAKEN_VILLAGES"`
    NameForsakenVillages            string  `json:"NAME_FORSAKEN_VILLAGES"`
    TroopMoveSpeedMultiply          float64 `json:"TROOP_MOVESPEED_MULTIPLY"`
    TroopProductionMultiply         float64 `json:"TROOP_PRODUCTION_MULTIPLY"`
    ResourceProductionMultiply      float64 `json:"RESOURCE_PRODUCTION_MULTIPLY"`
    BuildingConstructionTimeMultiply float64 `json:"BUILDING_CONSTRUCTION_TIME_MULTIPLY"`
    BuildingConstructionCostMultiply float64 `json:"BUILDING_CONSTRUCTION_COST_MULTIPLY"`
    WarehouseStorageMultiply        float64 `json:"WAREHOUSE_STORAGE_MULTIPLY"`
    DefenseBonusMultiply            float64 `json:"DEFENSE_BONUS_MULTIPLY"`
    PopulationBaseValue             int     `json:"POPULATION_BASE_VALUE"`
}

var Config *GameConfig

func LoadConfig() error {
    data, err := os.ReadFile("./config.json")
    if err != nil {
        return fmt.Errorf("could not read config file: %w", err)
    }

    if err := json.Unmarshal(data, &Config); err != nil {
        return fmt.Errorf("could not parse config: %w", err)
    }

    return nil
}