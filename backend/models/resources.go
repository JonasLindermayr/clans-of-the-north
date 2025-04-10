package models

type Resources struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	VillageID uint 	  `json:"villageID" gorm:"uniqueIndex`
	Wood      int    `json:"wood"`
	Stone     int    `json:"stone"`
	Clay	  int 	  `json:"clay"`
	Grain 	  int 	  `json:"grain"`

}
