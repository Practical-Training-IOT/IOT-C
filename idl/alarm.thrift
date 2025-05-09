namespace go iot.alarm

struct AlarmListReq{
    1: i32 page;
    2: i32 size;
}

struct AlarmListResp{
    1: list<AlarmList> list;
}

struct AlarmList{
    1: i32 id;
    2: string title;
    3: bool enabled;
    4: string type;
    5: string level;
    6: string desc;
    7: string time;
}

struct AlarmDetailReq{
    1: i32 id;
}

struct AlarmDetailResp{
    1: string name;
    2: string type;
    3: string status;
    4: string level;
    5: string time;
    6: string desc;
    7: string triggers;
    8: string notifies;
    9: string silence;
}

struct AlarmSearchReq{
    1: string title;
}

struct AlarmSearchResp{
    1: list<AlarmList> list;
}

service alarm {
    AlarmListResp AlarmList(1: AlarmListReq req);
    AlarmDetailResp AlarmDetail(1: AlarmDetailReq req);
    AlarmSearchResp AlarmSearch(1: AlarmSearchReq req);
}