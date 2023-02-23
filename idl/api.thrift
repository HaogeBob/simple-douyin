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

struct Video {
    1: i64 id      //视频唯一标识
	2: User author           //视频作者信息
	3: string play_url      //视频播放地址
	4: string cover_url     //视频封面地址
	5: i64 favorite_count   //视频的点赞总数
	6: i64 comment_count    //视频的评论总数
	7: bool is_favorite     //true-已点赞，false-未点赞
	8: string title       //视频标题
}

struct BaseResp {
	1: i32 status_code            //状态码
	2: string status_message   //状态描述
	3: i64 service_time         //服务时间
}

struct douyin_feed_request {
	1: i64 latest_time     //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	2: string token        //可选参数，登陆用户设置
}

struct fouyin_feed_response {
	1: BaseResp base_resp  
	2: list<Video> video_list   //视频列表
	3: i64 next_time               //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}


service ApiService {
    douyin_user_register_response DouyinUserRegister(1: douyin_user_register_request req) (api.post="/douyin/user/register/")
    douyin_user_login_response DouyinUserLogin(1: douyin_user_login_request req) (api.post="/douyin/user/login/")
    douyin_user_response DouyinUserGet(1: douyin_user_request req) (api.get="/douyin/user/")

    fouyin_feed_response Feed (1: douyin_feed_request req) (api.get="/douyin/feed/") 
}