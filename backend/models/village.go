package models

type Village struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	UserID    uint      `json:"userId"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	CordX    int       `json:"cordX"`
	CordY    int       `json:"cordY"`
	Points    int       `json:"points"`

	Resources Resources `json:"resources" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
