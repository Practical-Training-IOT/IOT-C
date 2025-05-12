package database

import "time"

type UserC struct {
	ID                uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username          string     `gorm:"column:username;size:50;unique;not null" json:"username"`
	Password          string     `gorm:"column:password;size:255;not null" json:"password,omitempty"` // 注意：通常不应直接暴露密码字段
	Email             string     `gorm:"column:email;size:100;unique;not null" json:"email"`
	PhoneNumber       string     `gorm:"column:phone_number;size:15" json:"phone_number,omitempty"`
	FirstName         string     `gorm:"column:first_name;size:50" json:"first_name,omitempty"`
	LastName          string     `gorm:"column:last_name;size:50" json:"last_name,omitempty"`
	RegistrationDate  time.Time  `gorm:"column:registration_date;default:CURRENT_TIMESTAMP" json:"registration_date"`
	LastLogin         *time.Time `gorm:"column:last_login" json:"last_login,omitempty"`
	ProfilePictureURL string     `gorm:"column:profile_picture_url" json:"profile_picture_url,omitempty"`
	Bio               string     `gorm:"column:bio" json:"bio,omitempty"`
	Department        string     `gorm:"column:department" json:"department,omitempty"`
}

// TableName 返回数据库中的实际表名
func (UserC) TableName() string {
	return "users_c"
}
