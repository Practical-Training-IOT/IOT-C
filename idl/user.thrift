namespace go iot.user

include "base.thrift"

struct RegisterReq {
    1: string userName
    2: string password
}

struct RegisterRes {
    1: i32 id
}

struct LoginReq {
    1: string userName
    2: string password
}

struct LoginRes {
    1: i32 id
}

struct UserInfoReq {
    1: i32 id
}

struct UserInfoRes {
    1: string username
    2: string department
    3: string email
    4: string phone
    5: string avatar
}

struct UserInfoUploadReq {
    1: string Department,
    2: string Email,
    3: i64 Phone,
    4: string Username
    5:i32 id
}

struct UserInfoUploadRes {
}

service user{
    RegisterRes Register(1:RegisterReq req)
    LoginRes Login(1:LoginReq req)
    UserInfoRes UserInfo(1:UserInfoReq req)
    UserInfoUploadRes UserInfoUpload(1:UserInfoUploadReq req)
}
