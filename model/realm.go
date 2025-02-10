package model

type Realms struct {
	ID       int `gorm:"primaryKey"`
	AdminID  int
	Admin    User            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Elements []RealmElements `gorm:"foreignKey:RealmID;"`
	Blocked  string
	SpawnX   int
	SpawnY   int
}

type Elements struct {
	ID         int `gorm:"primaryKey"`
	Animated   bool
	Filename   string
	HeightBlox int
	WidthBlox  int
	Frames     int
	Realms     []RealmElements `gorm:"foreignKey:ElementID;"`
}

type RealmElements struct {
	ID        int `gorm:"primaryKey"`
	X         int
	Y         int
	RealmID   int
	Realm     Realms `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ElementID int
	Element   Elements `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
