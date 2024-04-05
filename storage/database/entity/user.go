package entity

type User struct {
	Base
	Email          string `gorm:"unique;not null"`
	Username       string `gorm:"not null"`
	ProfileImgPath string
	Groups         []*Group `gorm:"many2many:group_members; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
