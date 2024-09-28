package entity

type User struct {
	Login    string `gorm:"primaryKey" json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}
