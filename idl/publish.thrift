namespace go publish

struct BaseResp {
    1: i32 status_code
    2: string status_message
}

struct User {
    1:i64 id
    2:string name
    3:optional i64 follow_count
    4:optional i64 follower_count
    5:bool is_follow
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

struct PublishActionRequest {
    1:i64 user_id
    2:string title
    3:string play_url
    4:string cover_url
}

struct PublishActionResponse {
    1:BaseResp baseResp
}

struct PublishListRequest {
    1:i64 user_id
    2:i64 now_user_id
}

struct PublishListResponse {
    1:BaseResp baseResp
    2:list<Video> video_list
}

service PublishService {
    PublishActionResponse PublishAction(1:PublishActionRequest req)
    PublishListResponse PublishList(1:PublishListRequest req)
}
