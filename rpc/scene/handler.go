package main

import (
	"context"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/config"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/database"
	scene "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/scene"
)

// SceneImpl implements the last service interface defined in the IDL.
type SceneImpl struct{}

// List implements the SceneImpl interface.
func (s *SceneImpl) List(ctx context.Context, req *scene.SceneListReq) (resp *scene.SceneListRes, err error) {
	// TODO: Your code here...
	var scenes []database.Scene
	err = config.DB.Where("deleted_at IS NULL").Find(&scenes).Error
	if err != nil {
		return nil, err
	}
	var sli []*scene.SceneList
	for _, i := range scenes {
		one := scene.SceneList{
			Id:      int32(i.ID),
			Title:   i.SceneName,
			Enabled: i.EnabledStatus,
			Desc:    i.SceneDescription,
			Time:    i.CreatedAt.String(),
		}
		sli = append(sli, &one)
	}
	return &scene.SceneListRes{List: sli}, nil
}

// UpdateEnable implements the SceneImpl interface.
func (s *SceneImpl) UpdateEnable(ctx context.Context, req *scene.SceneUpdateEnableReq) (resp *scene.SceneUpdateEnableRes, err error) {
	// TODO: Your code here...
	err = config.DB.Model(&database.Scene{}).Where("id=?", req.Id).Update("enabledstatus", req.Enable).Error
	if err != nil {
		return nil, err
	}
	return &scene.SceneUpdateEnableRes{}, nil
}

// Detail implements the SceneImpl interface.
func (s *SceneImpl) Detail(ctx context.Context, req *scene.SceneDetailReq) (resp *scene.SceneDetailRes, err error) {
	// TODO: Your code here...
	var scenes database.Scene
	var trigger []database.TriggerCondition
	var execution []database.ExecutionAction
	config.DB.Where("id=?", req.Id).First(&scenes)
	config.DB.Where("sceneid=?", req.Id).Find(&trigger)
	config.DB.Where("sceneid=?", req.Id).Find(&execution)
	var triSli []*scene.Trigger
	for _, i := range trigger {
		var event string
		switch i.TriggerMode {
		case "DeviceTriggered":
			event = "设备触发"
		case "TimeTriggered":
			event = "定时触发"
		default:
			event = "设备触发"
		}
		one := &scene.Trigger{
			Mode:      i.TriggerMode,
			Event:     event,
			Period:    "",
			Product:   i.Product,
			Device:    i.Device,
			Func:      i.Function,
			Condition: i.JudgmentCondition,
		}
		triSli = append(triSli, one)
	}
	var executionSli []*scene.Action
	for _, i := range execution {
		one := &scene.Action{
			Type:    i.ActionType,
			Product: i.Product,
			Device:  i.Device,
			Attr:    i.Function,
			Value:   i.Value,
		}
		executionSli = append(executionSli, one)
	}
	var status string
	switch scenes.EnabledStatus {
	case true:
		status = "启用"
	case false:
		status = "禁用"
	default:
		status = "启用"
	}
	return &scene.SceneDetailRes{
		Name:     scenes.SceneName,
		Status:   status,
		Time:     scenes.CreatedAt.String(),
		Desc:     scenes.SceneDescription,
		Triggers: triSli,
		Actions:  executionSli,
	}, nil
}
