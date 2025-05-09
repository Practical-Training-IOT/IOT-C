namespace go iot.user

include "base.thrift"

struct RegisterReq {
    1: string userName
}

struct RegisterRes {
    1: i32 id
    255: base.BaseResp BaseResp
}

service user{
    RegisterRes Register(1:RegisterReq req)
}
