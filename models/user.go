package models

type User struct {
	// gorm.Model
	Id        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"userName" gorm:"unique ;column:username"  form:"username" binding:"required" `
	FirstName string `json:"firstName" form:"firstName" gorm:"column:first_name" binding:"required"`
	LastName  string `json:"lastName" form:"lastName" gorm:"column:last_name" binding:"required"`
	Password  string `json:"-" form:"password" gorm:"column:password" binding:"required"`
	Email     string `json:"email" gorm:"unique ;column:email" form:"email"  binding:"required"`
	IsAdmin   bool   `json:"is_admin" gorm:"default:false"`
}
