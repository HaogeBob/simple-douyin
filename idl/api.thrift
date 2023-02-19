namespace go api

// struct BaseResp {
//     1: i64 status_code
//     2: string status_msg
// }

// 用户注册接口
struct douyin_user_register_request {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct douyin_user_register_response {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

// 用户登录接口
struct douyin_user_login_request {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct douyin_user_login_response {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

// 用户信息
struct douyin_user_request {
    1: i64 user_id (api.query="user_id")
    2: string token (api.query="token", api.vd="len($) > 0")
}

struct douyin_user_response {
    1: i32 status_code
    2: string status_msg
    3: User user
}

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

service ApiService {
    douyin_user_register_response DouyinUserRegister(1: douyin_user_register_request req) (api.post="/douyin/user/register/")
    douyin_user_login_response DouyinUserLogin(1: douyin_user_login_request req) (api.post="/douyin/user/login/")
    douyin_user_response DouyinUserGet(1: douyin_user_request req) (api.get="/douyin/user/")
}