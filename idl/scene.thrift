namespace go iot.scene

struct SceneListReq {

}

struct SceneList {
    1:i32 id
    2:string title
    3:bool enabled
    4:string desc
    5:string time
}

struct SceneListRes {
    1: list<SceneList> list
}

struct SceneUpdateEnableReq {
    1:i32 id
    2:bool enable
}

struct SceneUpdateEnableRes {

}

struct Trigger {
  1: string mode,
  2: string event,
  3: string period,
  4: string product,
  5: string device,
  6: string func,
  7: string condition
}

struct Action {
  1: string type,
  2: string product,
  3: string device,
  4: string attr,
  5: string value
}

struct SceneDetailRes {
  1: string name,
  2: string status,
  3: string time,
  4: string desc,
  5: list<Trigger> triggers,
  6: list<Action> actions
}

struct SceneDetailReq{
  1:i32 id
}

service scene{
    SceneListRes List(1:SceneListReq req)
    SceneUpdateEnableRes UpdateEnable(1:SceneUpdateEnableReq req)
    SceneDetailRes Detail(1:SceneDetailReq req)
}
