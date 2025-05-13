namespace go iot.ai

// 定义聊天响应结构
struct ChatResponse {
    1:string message,
    2:string model,
    3:i64 id,
}

struct ChatRequest {
    1:i32 id
    2:string message
    3:i32 userId
}

// 定义获取单个历史记录请求结构
struct OneHistoryRequest {
    1: i32 id,
    2: i32 userId
}

// 定义单个历史记录的聊天响应结构
struct ChatOneResponse {
    1:string message,
    2:string req,
    3:string model,
    4:i64 id,
}

// 定义单个历史记录响应结构
struct OneHistoryResponse {
    1:list<ChatOneResponse> chat,
    2:i64 id,
}

struct ChangeRequest {
    1:string message
}

struct ChangeResponse{

}

struct History{
    1:i32 id
    2:string title
    3:string created_at
}

struct HistoryResponse{
    1:list<History> list
}

struct HistoryRequest{
    1:i32 userId
}

service ai {
    OneHistoryResponse OneHistory(1:OneHistoryRequest req)
    ChatResponse Chat(1:ChatRequest req)
    ChangeResponse Change(1:ChangeRequest req)
    HistoryResponse History(1:HistoryRequest req)
}
