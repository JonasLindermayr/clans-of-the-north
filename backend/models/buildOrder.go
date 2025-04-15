package models

import "time"

type BuildOrder struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	VillageID uint   `json:"villageID"`
	BuildingID uint   `json:"buildingID"`
	StartTime time.Time `json:"startTime"`
	EndTime time.Time `json:"endTime"`
}