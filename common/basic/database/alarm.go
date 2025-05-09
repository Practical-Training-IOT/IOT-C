package database

import (
	"time"
)

// Alarm 对应数据库中的 alarms 表
type Alarm struct {
	ID               int64      `gorm:"primaryKey;autoIncrement" json:"id"`       // id
	RuleName         string     `gorm:"type:text" json:"rule_name"`               // 规则名称
	AlarmType        string     `gorm:"type:text" json:"alarm_type"`              // 告警类型
	AlarmLevel       string     `gorm:"type:text" json:"alarm_level"`             // 告警级别
	RuleDescription  string     `gorm:"type:text" json:"rule_description"`        // 规则描述
	TriggerMode      string     `gorm:"type:text" json:"trigger_mode"`            // 触发方式
	ProductID        int32      `gorm:"type:integer" json:"product_id"`           // 产品id
	EquipmentID      int32      `gorm:"type:integer" json:"equipment_id"`         // 设备id
	Function         string     `gorm:"size:50" json:"function"`                  // 功能
	ValueType        string     `gorm:"type:text" json:"value_type"`              // 取值类型
	ValuePeriod      string     `gorm:"type:text" json:"value_period"`            // 取值周期
	JudgingCondition string     `gorm:"type:text" json:"judging_condition"`       // 判断条件
	Value            string     `gorm:"type:text" json:"value"`                   // 取值
	SilencePeriod    string     `gorm:"type:text" json:"silence_period"`          // 静默时间
	MeanNotification string     `gorm:"type:text" json:"mean_notification"`       // 通知方式
	Status           string     `gorm:"type:text" json:"status"`                  // 状态
	CreatedAt        time.Time  `gorm:"type:timestamptz" json:"created_at"`       // 创建时间
	UpdatedAt        time.Time  `gorm:"type:timestamptz" json:"updated_at"`       // 更新时间
	DeletedAt        *time.Time `gorm:"type:timestamptz;index" json:"deleted_at"` // 删除时间
	CreatedBy        int64      `gorm:"type:bigint" json:"created_by"`            // 创建者
	UpdatedBy        int64      `gorm:"type:bigint" json:"updated_by"`            // 更新者
	DeletedBy        *int64     `gorm:"type:bigint" json:"deleted_by"`            // 删除者
}

// TableName 设置表名
func (Alarm) TableName() string {
	return "alarms"
}
