package models

type Tile struct {
	ID        uint `gorm:"primaryKey"`
	X         int  `gorm:"not null;index:idx_coords,unique"`
	Y         int  `gorm:"not null;index:idx_coords,unique"`
	VillageID *uint	`json:"villageID"`
	Terrain   string `gorm:"default:plains"`
}