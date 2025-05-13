package database

import "time"

type Ai struct {
	ID       int64  `gorm:"primaryKey" json:"id"`                // 主键id
	Model    string `gorm:"column:model;size:50" json:"model"`   // 使用的哪个ai
	Req      string `gorm:"column:req;type:text" json:"req"`     // 发送的请求
	Res      string `gorm:"column:res;type:text" json:"res"`     // 回答的话语
	Title    string `gorm:"column:title;type:text" json:"title"` // 标题
	AiScenId int64  `gorm:"column:ai_scen_id" json:"ai_scen_id"`
	UserId   int64  `gorm:"column:user_id" json:"user_id"`
}

// TableName 指定数据库表名
func (Ai) TableName() string {
	return "ai"
}

type AiScene struct {
	ID        int64     `gorm:"primaryKey" json:"id"`                // 主键id
	Title     string    `gorm:"column:title;type:text" json:"title"` // 场景标题
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
	UserId    int64     `gorm:"column:user_id" json:"user_id"`
}

// TableName 指定数据库表名
func (AiScene) TableName() string {
	return "ai_scene"
}
