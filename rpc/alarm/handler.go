package main

import (
	"context"
	alarm "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/alarm"
)

// AlarmImpl implements the last service interface defined in the IDL.
type AlarmImpl struct{}

// AlarmList implements the AlarmImpl interface.
func (s *AlarmImpl) AlarmList(ctx context.Context, req *alarm.AlarmListReq) (resp *alarm.AlarmListResp, err error) {
	// TODO: Your code here...
	
	return
}

// AlarmDetail implements the AlarmImpl interface.
func (s *AlarmImpl) AlarmDetail(ctx context.Context, req *alarm.AlarmDetailReq) (resp *alarm.AlarmDetailResp, err error) {
	// TODO: Your code here...
	return
}

// AlarmSearch implements the AlarmImpl interface.
func (s *AlarmImpl) AlarmSearch(ctx context.Context, req *alarm.AlarmSearchReq) (resp *alarm.AlarmSearchResp, err error) {
	// TODO: Your code here...
	return
}
