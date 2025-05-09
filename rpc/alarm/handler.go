package main

import (
	"context"
	"fmt"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/config"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/database"
	alarm "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/alarm"
)

// AlarmImpl implements the last service interface defined in the IDL.
type AlarmImpl struct{}

// AlarmList implements the AlarmImpl interface.
func (s *AlarmImpl) AlarmList(ctx context.Context, req *alarm.AlarmListReq) (resp *alarm.AlarmListResp, err error) {
	// TODO: Your code here...
	var AlarmList []*alarm.AlarmList
	var total int64
	var List []database.Alarm
	// 先获取总数
	err = config.DB.Model(&database.Alarm{}).Where("deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return &alarm.AlarmListResp{}, err
	}

	// 计算偏移量
	offset := (req.Page - 1) * req.Size

	// 查询分页数据
	err = config.DB.Where("deleted_at IS NULL").
		Order("created_at DESC").
		Offset(int(offset)).
		Limit(int(req.Size)).
		Find(&List).Error

	if err != nil {
		return &alarm.AlarmListResp{}, err
	}

	for _, v := range List {
		sprintf := fmt.Sprintf("%v", v.CreatedAt)
		if v.Status == "Enable" {
			list := alarm.AlarmList{
				Id:      int32(v.ID),
				Title:   v.RuleName,
				Enabled: false,
				Type:    v.AlarmType,
				Level:   v.AlarmLevel,
				Desc:    v.RuleDescription,
				Time:    sprintf,
			}
			AlarmList = append(AlarmList, &list)
		} else {
			list := alarm.AlarmList{
				Id:      int32(v.ID),
				Title:   v.RuleName,
				Enabled: true,
				Type:    v.AlarmType,
				Level:   v.AlarmLevel,
				Desc:    v.RuleDescription,
				Time:    sprintf,
			}
			AlarmList = append(AlarmList, &list)
		}
	}
	return &alarm.AlarmListResp{List: AlarmList}, nil
}

// AlarmDetail implements the AlarmImpl interface.
func (s *AlarmImpl) AlarmDetail(ctx context.Context, req *alarm.AlarmDetailReq) (resp *alarm.AlarmDetailResp, err error) {
	// TODO: Your code here...
	var alarms database.Alarm
	err = config.DB.Where("id = ?", req.Id).First(&alarms).Error
	if err != nil {
		return &alarm.AlarmDetailResp{}, err
	}
	time := fmt.Sprintf("%v", alarms.CreatedAt)
	return &alarm.AlarmDetailResp{
		Name:     alarms.RuleName,
		Type:     alarms.AlarmType,
		Status:   alarms.Status,
		Level:    alarms.AlarmLevel,
		Time:     time,
		Desc:     alarms.RuleDescription,
		Triggers: alarms.TriggerMode,
		Notifies: alarms.MeanNotification,
		Silence:  alarms.SilencePeriod,
	}, nil
}

// AlarmSearch implements the AlarmImpl interface.
func (s *AlarmImpl) AlarmSearch(ctx context.Context, req *alarm.AlarmSearchReq) (resp *alarm.AlarmSearchResp, err error) {
	// TODO: Your code here...
	var AlarmList []*alarm.AlarmList
	var List []database.Alarm
	err = config.DB.Model(&database.Alarm{}).Where("rule_name like ?", "%"+req.Title+"%").Find(&List).Error
	if err != nil {
		return &alarm.AlarmSearchResp{}, err
	}
	for _, v := range List {
		sprintf := fmt.Sprintf("%v", v.CreatedAt)
		if v.Status == "Enable" {
			list := alarm.AlarmList{
				Id:      int32(v.ID),
				Title:   v.RuleName,
				Enabled: false,
				Type:    v.AlarmType,
				Level:   v.AlarmLevel,
				Desc:    v.RuleDescription,
				Time:    sprintf,
			}
			AlarmList = append(AlarmList, &list)
		} else {
			list := alarm.AlarmList{
				Id:      int32(v.ID),
				Title:   v.RuleName,
				Enabled: true,
				Type:    v.AlarmType,
				Level:   v.AlarmLevel,
				Desc:    v.RuleDescription,
				Time:    sprintf,
			}
			AlarmList = append(AlarmList, &list)
		}
	}
	return &alarm.AlarmSearchResp{List: AlarmList}, nil
}
