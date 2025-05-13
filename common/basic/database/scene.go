package database

import (
	"database/sql"
	"time"
)

type ExecutionAction struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SceneID     *int   `gorm:"column:sceneid" json:"scene_id,omitempty"` // 外键，关联 scenes.id
	ActionOrder int    `gorm:"column:actionorder;not null" json:"action_order"`
	ActionType  string `gorm:"column:actiontype;type:varchar(50);not null" json:"action_type"`
	Product     string `gorm:"column:product;type:varchar(255)" json:"product,omitempty"`
	Device      string `gorm:"column:device;type:varchar(255)" json:"device,omitempty"`
	Function    string `gorm:"column:function;type:varchar(255)" json:"function,omitempty"`
	Value       string `gorm:"column:value;type:varchar(255)" json:"value,omitempty"`
}

// TableName 设置对应的数据库表名
func (ExecutionAction) TableName() string {
	return "executionactions"
}

type Scene struct {
	ID               int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SceneName        string     `gorm:"column:scenename;type:text" json:"scene_name,omitempty"`
	SceneDescription string     `gorm:"column:scenedescription;type:text" json:"scene_description,omitempty"`
	CreationTime     *time.Time `gorm:"column:creationtime;type:timestamptz(6)" json:"creation_time,omitempty"`
	EnabledStatus    bool       `gorm:"column:enabledstatus" json:"enabled_status"`
	CreatedAt        *time.Time `gorm:"column:created_at;type:timestamptz(6)" json:"created_at,omitempty"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:timestamptz(6)" json:"updated_at,omitempty"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;index:idx_scenes_deleted_at;type:timestamptz(6)" json:"deleted_at,omitempty"`
}

func (Scene) TableName() string {
	return "scenes"
}

type TriggerCondition struct {
	ID                int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SceneID           *int           `gorm:"column:sceneid;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"scene_id,omitempty"` // 外键，关联 scenes.id
	TriggerType       string         `gorm:"column:triggertype;type:varchar(50);not null" json:"trigger_type"`
	Product           string         `gorm:"column:product;type:varchar(255)" json:"product,omitempty"`
	Device            string         `gorm:"column:device;type:varchar(255)" json:"device,omitempty"`
	Function          string         `gorm:"column:function;type:varchar(255)" json:"function,omitempty"`
	ValueType         string         `gorm:"column:valuetype;type:varchar(50)" json:"value_type,omitempty"`
	JudgmentCondition string         `gorm:"column:judgmentcondition;type:varchar(255)" json:"judgment_condition,omitempty"`
	Time              sql.NullString `gorm:"column:time;type:time(6)" json:"time,omitempty"` // 使用 NullString 支持 NULL 值
	DaysOfWeek        string         `gorm:"column:daysofweek;type:varchar(50)" json:"days_of_week,omitempty"`
	TriggerMode       string         `gorm:"column:triggermode;type:varchar(50)" json:"trigger_mode,omitempty"`
}

func (TriggerCondition) TableName() string {
	return "triggerconditions"
}
