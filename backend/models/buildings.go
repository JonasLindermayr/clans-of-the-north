package models

type Buildings struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	VillageID uint   `json:"villageID"`
	BuildingID int
	Type      string
	CurrentLevel     int
	//Village   Village `gorm:"foreignKey:VillageID"`
}