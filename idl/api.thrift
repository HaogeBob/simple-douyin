namespace go api

struct douyin_publish_action_request {
    1:string token(api.form = 'token', api.vd="len($) > 0")
    2:list<byte> data(api.form = 'data', api.vd="len($) > 0")
    3:string title(api.form = 'title', api.vd="len($) > 0")
}

struct douyin_publish_action_response {
    1: i32 status_code
    2: string status_msg
}

struct douyin_publish_list_request {
    1: i64 user_id (api.query="user_id")
    2: string token (api.query="token", api.vd="len($) > 0")
}

struct douyin_publish_list_response {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
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
    1:i64 id
    2:User user
    3:string play_url
    4:string cover_url
    5:i64 favorite_count
    6:i64 comment_count
    7:bool is_favorite
    8:string title
}

service ApiService {
    douyin_publish_action_response DouyinPublishAction(1: douyin_publish_action_request req) (api.post="/douyin/publish/action/")
    douyin_publish_list_response DouyinPublishList(1: douyin_publish_list_request req) (api.post="/douyin/publish/list/")
}