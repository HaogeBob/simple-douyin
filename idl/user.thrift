namespace go user

// struct BaseResp {
//     1: i64 status_code
//     2: string status_message
//     3: i64 service_time
// }

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: string background_image
    8: string signature
    9: i64 total_favorited
    10: i64 work_count
    11: i64 favorite_count
}

struct CreateUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct CreateUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
}

struct CheckUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct CheckUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
}

struct MGetUserRequest {
    1: i64 user_id
}

struct MGetUserResponse {
    1: i32 status_code
    2: string status_msg
    3: User user
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    CheckUserResponse CheckUser(1: CheckUserRequest req)
    MGetUserResponse MGetUser(1: MGetUserRequest req)
}