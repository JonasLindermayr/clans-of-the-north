package models

import "time"

type Resources struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	VillageID uint 	  `json:"villageID" gorm:"uniqueIndex"`
	Wood      int    `json:"wood"`
	Stone     int    `json:"stone"`
	Clay	  int 	  `json:"clay"`
	Grain 	  int 	  `json:"grain"`
	Gold      int    `json:"gold"`
	ClaimedPopulation int    `json:"claimedPopulation"`
	Population int    `json:"population"`
	Storage   int    `json:"storage"`
	CarryOverWood float64 `json:"carryOverWood"`
	CarryOverStone float64 `json:"carryOverStone"`
	CarryOverClay float64 `json:"carryOverClay"`
	CarryOverGrain float64 `json:"carryOverGrain"`
	CarryOverGold float64 `json:"carryOverGold"`
	LastResourceUpdate time.Time  `json:"lastResourceUpdate"`
}
