package model

type User struct {
	ID			int		`gorm:"primaryKey"`
	Name 		string
	Username 	string		`gorm:"unique"`
	Email 		string		`gorm:"unique"`
	AvatarID	uint
	Avatar		Avatar		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Realms		[]Realms		`gorm:"foreignKey:AdminID;"`
}

type Avatar struct {
	ID		int	`gorm:"primaryKey"`
	Users	[]User	`gorm:"foreignKey:AvatarID;"`
}