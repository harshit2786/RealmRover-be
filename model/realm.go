package model

type Realms struct {
	ID       uint 				`gorm:"primaryKey"`
	AdminID  uint
	Admin    User             	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Elements []RealmElements 	`gorm:"foreignKey:RealmID;"`
}

type Elements struct {
	ID     uint 				`gorm:"primaryKey"`
	Height int
	Width  int
	Realms  []RealmElements 	`gorm:"foreignKey:ElementID;"`
}

type RealmElements struct {
	ID        uint 		`gorm:"primaryKey"`
	X         int
	Y         int
	RealmID   uint
	Realm     Realms    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ElementID uint     
	Element   Elements `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
