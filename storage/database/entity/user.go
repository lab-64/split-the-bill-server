package entity

type User struct {
	Base
	Email    string `gorm:"unique;not null"`
	Username string
	Groups   []*Group `gorm:"many2many:group_members; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
